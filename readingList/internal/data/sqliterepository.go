package data

import (
	"context"
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

const (
	SqliteInsertBookCommand           = "INSERT INTO Books (title, published, pages, rating) VALUES (?, ?, ?, ?) RETURNING id, created_at, version"
	SqliteInsertGenreCommand          = "INSERT INTO Genres (book_id, genre_name) VALUES (?, ?) RETURNING id"
	SqliteSelectBookCommand           = "SELECT id, created_at, title, published, pages, rating, version FROM Books WHERE id = ?"
	SqliteSelectGenresByBookIdCommand = "SELECT genre_name FROM Genres Where book_id = ?"
	SqliteDeleteBookCommand           = "DELETE FROM Books WHERE id = ?"
)

var (
	NotFoundError = errors.New("error")
)

type SqliteRepository struct {
	db *sql.DB
}

func NewSqliteRepository(dsn string) (Repository, error) {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}

	return &SqliteRepository{db: db}, nil
}

func (r *SqliteRepository) Migrate() error {
	createBooks := `
	CREATE TABLE IF NOT EXISTS Books (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
		title TEXT NOT NULL,
	    published INTEGER NOT NULL,
	    pages INTEGER NOT NULL,
	    rating REAL NOT NULL,
	    version INTEGER NOT NULL default 1
	);
	`
	createGenresByBooks := `
	CREATE TABLE IF NOT EXISTS Genres (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
	    book_id INTEGER NOT NULL,
	    genre_name text NOT NULL,
	    CONSTRAINT FK_BOOKS 
	        FOREIGN KEY (book_id) 
			REFERENCES Books(id)
			ON DELETE CASCADE 	                                  
	);
	`

	tx, err := r.db.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}

	if _, err := tx.Exec(createBooks); err != nil {
		return err
	}
	if _, err := tx.Exec(createGenresByBooks); err != nil {
		return err
	}

	return tx.Commit()
}

func (r *SqliteRepository) Ping() error {
	return r.db.Ping()
}

func (r *SqliteRepository) Close() error {
	return r.db.Close()
}

func (r *SqliteRepository) InsertOne(title string, published int, pages int, rating float64, genres []string) (*Book, error) {
	args := []interface{}{title, published, pages, rating}
	var bookId int64
	var createdAt int64
	var version int64
	err := r.db.QueryRow(SqliteInsertBookCommand, args...).
		Scan(&bookId, &createdAt, &version)
	if err != nil {
		return nil, err
	}

	for _, genre := range genres {
		if _, err := r.db.Exec(SqliteInsertGenreCommand, bookId, genre); err != nil {
			if deleteErr := r.DeleteById(bookId); deleteErr != nil {
				return nil, errors.Join(err, deleteErr)
			} else {
				return nil, err
			}
		}
	}

	book, err := r.FindById(bookId, true)
	return book, err
}

func (r *SqliteRepository) FindById(id int64, includeGenres bool) (*Book, error) {
	if id < 1 {
		return nil, NotFoundError
	}

	var book Book
	err := r.db.QueryRow(SqliteSelectBookCommand).Scan(
		&book.ID,
		&book.CreatedAt,
		&book.Title,
		&book.Published,
		&book.Pages,
		&book.Rating,
		&book.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, NotFoundError
		default:
			return nil, err
		}
	}

	if includeGenres {
		genres, err := r.LoadGenres(&book)
		if err != nil {
			return nil, err
		}
		book.Genres = genres
	}

	// obviously wrong but leaving it for now
	return &book, nil
}

func (r *SqliteRepository) LoadGenres(book *Book) ([]string, error) {
	if book == nil {
		return nil, NotFoundError
	}

	rows, err := r.db.Query(SqliteSelectGenresByBookIdCommand, book.ID)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Print(err)
		}
	}()

	var genres []string
	for rows.Next() {
		var genre string
		if err := rows.Scan(&genre); err != nil {
			return nil, err
		}
		genres = append(genres, genre)
	}

	return genres, nil
}

func (r *SqliteRepository) DeleteById(id int64) error {
	if id < 1 {
		return NotFoundError
	}
	result, err := r.db.Exec(SqliteDeleteBookCommand, id)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return NotFoundError
		default:
			return err
		}
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected < 1 {
		return NotFoundError
	}
	return nil
}
