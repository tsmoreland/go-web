package data

import (
	"context"
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
)

const (
	SqliteInsertBookCommand           = "INSERT INTO Books (title, published, pages, rating) VALUES (?, ?, ?, ?) RETURNING id, created_at, version"
	SqliteInsertGenreCommand          = "INSERT INTO Genres (book_id, genre_name) VALUES (?, ?) RETURNING id"
	SqliteSelectBookCommand           = "SELECT id, created_at, title, published, pages, rating, version WHERE id = ?"
	SqliteSelectGenresByBookIdCommand = "SELECT book_id, genre_name FROM Genres Where book_id = ?"
	SqliteDeleteBookCommand           = "DELETE FROM Books WHERE id = ?"
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
			if _, deleteErr := r.DeleteById(bookId); deleteErr != nil {
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

	// obviously wrong but leaving it for now
	return nil, nil
}

func (r *SqliteRepository) DeleteById(id int64) (int64, error) {
	result, err := r.db.Exec(SqliteDeleteBookCommand, id)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	return rowsAffected, err
}
