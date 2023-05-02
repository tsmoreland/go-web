package main

import (
	"fmt"
	"github.com/tsmoreland/go-web/readingList/internal/models"
	"html/template"
	"net/http"
	"strconv"
	"strings"
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

	files := []string{
		"./ui/html/base.html",
		"./ui/html/partials/nav.html",
		"./ui/html/pages/home.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.logger.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", books)
	if err != nil {
		app.logger.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (app *application) bookView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	book, err := app.client.Get(int64(id))
	if err != nil {
		app.logger.Print(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	_, _ = fmt.Fprintf(w, "%s (%d)", book.Title, book.Pages)
}

func (app *application) bookCreate(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		app.bookCreateForm(w, r)
	case http.MethodPost:
		app.bookCreateProcess(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (app *application) bookCreateForm(w http.ResponseWriter, r *http.Request) {
	_ = r
	_, _ = fmt.Fprintf(w, `
		<html>
			<head><title>Create Book</title></head>
			<body>
				<h1>Create Book</h1>
				<form action="/book/create" method="POST">
					<label for="title">Title></label>
					<input type="text" name="title" id="title">
					<label for="pages">Pages></label>
					<input type="text" name="pages" id="pages">
					<label for="published">Published></label>
					<input type="text" name="published" id="published">
					<label for="genres">Genres></label>
					<input type="text" name="genres" id="genres">
					<label for="rating">Rating></label>
					<input type="text" name="rating" id="rating">
					<button type="submit">Create</button>
				</form>
			</body>
		</html>
		`)
}

func (app *application) bookCreateProcess(w http.ResponseWriter, r *http.Request) {
	title := r.PostFormValue("title")
	if title == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	pages, err := strconv.Atoi(r.PostFormValue("pages"))
	if err != nil || pages < 1 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	published, err := strconv.Atoi(r.PostFormValue("pages"))
	if err != nil || published < 1 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	genres := strings.Split(r.PostFormValue("genres"), " ")
	_ = genres

	rating, err := strconv.ParseFloat(r.PostFormValue("rating"), 64)
	if err != nil || rating < 0.0 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	addOrUpdateBook := &models.AddOrUpdateBook{
		Title:     title,
		Pages:     pages,
		Published: published,
		Genres:    genres,
		Rating:    rating,
	}

	if _, err = app.client.Create(addOrUpdateBook); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *application) bookEdit(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		app.bookEditForm(w, r)
	case http.MethodPost:
		app.bookEditProcess(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (app *application) bookEditForm(w http.ResponseWriter, r *http.Request) {
	_ = w
	_ = r
}

func (app *application) bookEditProcess(w http.ResponseWriter, r *http.Request) {
	_ = w
	_ = r
}
