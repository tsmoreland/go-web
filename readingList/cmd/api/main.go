package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type options struct {
	port int
	env  string
}

type service struct {
	settings options
	logger   *log.Logger
}

func main() {
	var serverOptions options

	flag.IntVar(&serverOptions.port, "port", 9000, "HTTP Listen Port")
	flag.StringVar(&serverOptions.env, "env", "production", "server environment")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	svc := &service{
		settings: serverOptions,
		logger:   logger,
	}
	serverAddress := fmt.Sprintf(":%d", svc.settings.port)

	svr := http.Server{
		Addr:         serverAddress,
		Handler:      svc.route(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("Starting %s server on %s", svc.settings.env, serverAddress)
	if err := svr.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
