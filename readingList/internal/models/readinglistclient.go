package models

import (
	"fmt"
	"log"
	"net/http"
)

type Book struct {
	ID        int64    `json:"id"`
	Title     string   `json:"title"`
	Published int      `json:"published"`
	Pages     int      `json:"pages"`
	Genres    []string `json:"geres"`
	Rating    float64  `json:"rating"`
}

type BookResponse struct {
	Book *Book `json:"book"`
}

type BooksResponse struct {
	Books *[]Book `json:"books"`
}

type ReadingListClient struct {
	Endpoint string
}

func (c *ReadingListClient) GetAll() (*[]Book, error) {
	resp, err := http.Get(c.Endpoint)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Println(err)
		}
	}()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}
	var booksResponse BooksResponse
	if err := ReadJSONObject(resp.Body, &booksResponse); err != nil {
		return nil, err
	}
	return booksResponse.Books, nil
}

func (c *ReadingListClient) Get(id int64) (*Book, error) {
	url := fmt.Sprintf("%s/%d", c.Endpoint, id)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Println(err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	var bookResponse BookResponse
	if err := ReadJSONObject(resp.Body, &bookResponse); err != nil {
		return nil, err
	}
	return bookResponse.Book, nil
}
