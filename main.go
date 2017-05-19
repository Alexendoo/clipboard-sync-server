package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/Alexendoo/sync/routes"
)

func main() {
	db, err := sql.Open("postgres", "postgres://devel:devel@localhost/sync?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	srv := &http.Server{
		Addr:         "localhost:8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      routes.Handler(db),
	}

	log.Fatal(srv.ListenAndServe())
}
