package main

import (
	"encoding/json"
	"fmt"
	"github.com/tsmoreland/go-web/readingList/internal/data"
	"net/http"
	"strconv"
	"time"

	_ "github.com/tsmoreland/go-web/readingList/internal/data"
)

func (svc *service) healthcheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	status := map[string]string{
		"status":      "available",
		"environment": svc.settings.env,
	}

	if js, err := json.Marshal(status); err == nil {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(js)
	} else {
		http.Error(w, "", http.StatusInternalServerError)
	}
}

func (svc *service) getOrCreateBooks(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		books := []data.Book{
			{
				ID:        1,
				CreatedAt: time.Date(1990, 11, 20, 0, 0, 0, 0, time.UTC),
				Title:     "Jurassic Park",
				Pages:     448,
				Genres:    []string{"Science Fiction", "Action"},
				Rating:    5.0,
				Version:   1.0,
			},
			{
				ID:        2,
				CreatedAt: time.Date(1994, 11, 20, 0, 0, 0, 0, time.UTC),
				Title:     "Jurassic Park 2",
				Pages:     448,
				Genres:    []string{"Science Fiction", "Action"},
				Rating:    5.0,
				Version:   1.0,
			},
		}

		if err := svc.writeJSON(w, http.StatusOK, envelope{"books": books}); err != nil {
			http.Error(w, "", http.StatusInternalServerError)
		}
		return
	}
	if r.Method == http.MethodPost {
		var bookDto struct {
			Title     string   `json:"title"`
			Published int      `json:"published"`
			Pages     int      `json:"pages"`
			Genres    []string `json:"genres"`
			Rating    float64  `json:"rating"`
		}

		if err := svc.readJSONObject(w, r, &bookDto); err != nil {
			// custom error handling details: Alex Edwards, Let's Go Further Chapter 4
			svc.writeBadRequest(w, err)
			return
		}
		_, _ = fmt.Fprint(w, bookDto)
		return
	}
	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
}

func (svc *service) getUpdateOrDeleteBooks(w http.ResponseWriter, r *http.Request) {

	getId := func(r *http.Request) (int64, error) {
		id, err := strconv.ParseInt(r.URL.Path[len("api/v1/books/"):], 10, 64)
		if err != nil {
			http.Error(w, "", http.StatusBadRequest)
		}
		return id, err
	}

	switch r.Method {
	case http.MethodGet:
		if id, err := getId(r); err == nil {
			svc.getBook(id, w, r)
		}
	case http.MethodPut:
		if id, err := getId(r); err == nil {
			svc.updateBook(id, w, r)
		}
	case http.MethodDelete:
		if id, err := getId(r); err == nil {
			svc.deleteBook(id, w, r)
		}
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (svc *service) getBook(id int64, w http.ResponseWriter, r *http.Request) {
	_ = r

	book := data.Book{
		ID:        id,
		CreatedAt: time.Date(1990, 11, 20, 0, 0, 0, 0, time.UTC),
		Title:     "Jurassic Park",
		Pages:     448,
		Genres:    []string{"Science Fiction", "Action"},
		Rating:    5.0,
		Version:   1.0,
	}

	if err := svc.writeJSON(w, http.StatusOK, envelope{"book": book}); err != nil {
		http.Error(w, "", http.StatusInternalServerError)
	}
}
func (svc *service) updateBook(id int64, w http.ResponseWriter, r *http.Request) {
	var bookDto struct {
		Title     *string  `json:"title"`
		Published *int     `json:"published"`
		Pages     *int     `json:"pages"`
		Genres    []string `json:"genres"`
		Rating    *float64 `json:"rating"`
	}

	book := data.Book{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Hunger Games",
		Published: 2008,
		Pages:     374,
		Rating:    5,
		Version:   1,
	}

	if err := svc.readJSONObject(w, r, &bookDto); err != nil {
		svc.writeBadRequest(w, err)
		return
	}

	if bookDto.Title != nil {
		book.Title = *bookDto.Title
	}
	if bookDto.Published != nil {
		book.Published = *bookDto.Published
	}
	if bookDto.Pages != nil {
		book.Pages = *bookDto.Pages
	}
	if len(bookDto.Genres) > 0 {
		book.Genres = bookDto.Genres
	}
	if bookDto.Rating != nil {
		book.Rating = *bookDto.Rating
	}

	_, _ = fmt.Fprintf(w, "update book %d %v", id, book)
}
func (svc *service) deleteBook(id int64, w http.ResponseWriter, r *http.Request) {
	_ = r
	_, _ = fmt.Fprintf(w, "delete book %d", id)
}
