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

		js, err := json.Marshal(books)
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(js)
		return
	}
	if r.Method == http.MethodPost {
		_, _ = fmt.Fprint(w, "")
		http.Error(w, "", http.StatusBadRequest)
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

	js, err := json.Marshal(book)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
}
func (svc *service) updateBook(id int64, w http.ResponseWriter, r *http.Request) {
	_ = r
	_, _ = fmt.Fprintf(w, "update book %d", id)
}
func (svc *service) deleteBook(id int64, w http.ResponseWriter, r *http.Request) {
	_ = r
	_, _ = fmt.Fprintf(w, "delete book %d", id)
}
