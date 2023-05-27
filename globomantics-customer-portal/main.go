package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/tsmoreland/go-web/globomantics-customer-portal/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const (
	hostPort = 8010
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/login", handlers.Login)
	r.HandleFunc("/logout", handlers.Logout)

	customerRouter := r.Methods(http.MethodGet).Subrouter()
	customerRouter.HandleFunc("/customers", handlers.Customers)

	r.PathPrefix("/home").Handler(http.FileServer(http.Dir("./public")))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/home", http.StatusMovedPermanently)
	})

	hostAddress := fmt.Sprintf(":%d", hostPort)
	server := &http.Server{
		Handler:      r,
		Addr:         hostAddress,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		log.Println("Starting server...")
		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Println(err)
	}
	log.Println("Server Shutting down")
	os.Exit(0)
}
