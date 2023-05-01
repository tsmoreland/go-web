package main

import (
	"fmt"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	books, err := app.client.GetAll()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	_, _ = fmt.Fprintf(w, "<html><head><title>Reading List</title></head><body><h1>Reading List</h1><ul>")

	for _, book := range *books {
		_, _ = fmt.Fprintf(w, "<li>%s (%d)</li>", book.Title, book.Pages)
	}

	_, _ = fmt.Fprintf(w, "</ul></body></html>")
	w.WriteHeader(http.StatusOK)
}

func (app *application) bookView(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (app *application) bookCreate(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (app *application) bookEdit(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
