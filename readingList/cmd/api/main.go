package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
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
	flag.StringVar(&serverOptions.env, "env", 9000, "server environment")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	svc := &service{
		settings: serverOptions,
		logger:   logger,
	}
	serverAddress := fmt.Sprintf(":%d", svc.settings.port)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/health", svc.healthcheck)

	logger.Printf("Starting %s server on %s", svc.settings.env, serverAddress)
	if err := http.ListenAndServe(serverAddress, mux); err != nil {
		log.Fatalln(err)
	}
}
