/*
 * readingList
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package main

import (
	"github.com/tsmoreland/go-web/readingList/internal/data"
	"log"
	"net/http"
	"os"
)

type Api struct {
	logger     *log.Logger
	repository *data.Repository
}

func NewApi(dsn string) (*Api, error) {
	repository, err := data.NewSqliteRepository(dsn)
	if err != nil {
		return nil, err
	}

	api := &Api{
		logger:     log.New(os.Stdout, "", log.Ldate|log.Ltime),
		repository: &repository,
	}
	return api, nil
}

func (api *Api) AddBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (api *Api) DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (api *Api) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (api *Api) GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (api *Api) UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
