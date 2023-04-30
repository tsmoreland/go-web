package main

import "net/http"

func (app *application) home(w http.ResponseWriter, r *http.Request) {
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
