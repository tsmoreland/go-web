package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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

	if bytes, err := json.Marshal(status); err == nil {
		_, _ = fmt.Fprintf(w, string(bytes))
	} else {
		http.Error(w, "", http.StatusInternalServerError)
	}
}

func (svc *service) getOrCreateBooks(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Print(w, "[]")
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
	_, _ = fmt.Fprintf(w, "Display book %d", id)
}
func (svc *service) updateBook(id int64, w http.ResponseWriter, r *http.Request) {
	_ = r
	_, _ = fmt.Fprintf(w, "update book %d", id)
}
func (svc *service) deleteBook(id int64, w http.ResponseWriter, r *http.Request) {
	_ = r
	_, _ = fmt.Fprintf(w, "delete book %d", id)
}
