package main

import (
	"log"
	"net/http"
)

func main() {
	log.Printf("Server started")

	api := NewApi()
	router := NewRouter(api)
	log.Fatal(http.ListenAndServe(":9000", router))
}
