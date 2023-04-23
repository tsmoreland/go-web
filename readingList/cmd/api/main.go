package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/api/v1/health", healthcheck)

	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{\"status\": \"available\"}")
}
