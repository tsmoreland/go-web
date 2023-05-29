package ordersApi

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tsmoreland/go-web/ordersApi/handlers"
)

const port = 4000

func main() {
	fmt.Println("Welcome to the Orders App!")
	handler, err := handlers.New()
	if err != nil {
		log.Fatal(err)
	}
	// start server
	router := handlers.ConfigureHandler(handler)
	fmt.Printf("Listening on localhost:%d...", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
