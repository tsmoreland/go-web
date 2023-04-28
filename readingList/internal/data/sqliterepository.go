package data

import (
	"context"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
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

func (r *SqliteRepository) Close() error {
	return r.db.Close()
}
