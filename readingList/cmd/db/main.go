package main

import (
	"flag"
	"github.com/tsmoreland/go-web/readingList/internal/data"
	"log"
)

func main() {

	var dsn string
	flag.StringVar(&dsn, "dsn", "readingList.db", "database service name")
	flag.Parse()

	r, err := data.NewSqliteRepository(dsn)
	if err != nil {
		log.Fatalln(err)
	}

	defer func() {
		if err := r.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if err := r.Migrate(); err != nil {
		log.Fatal(err)
	}
}
