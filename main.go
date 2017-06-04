package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/Alexendoo/clipboard-sync-server/routes"
)

func main() {
	db, err := sql.Open("postgres", "postgres://devel:devel@localhost/sync?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	srv := &http.Server{
		Addr:              "localhost:8080",
		ReadHeaderTimeout: 10 * time.Second,
		IdleTimeout:       2 * time.Minute,
		Handler:           routes.Handler(db),
	}

	log.Fatal(srv.ListenAndServe())
}
