package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/Alexendoo/clipboard-sync-server/config"
	"github.com/Alexendoo/clipboard-sync-server/routes"
)

var (
	configDir  = flag.String("configdir", config.DefaultDir(), "Use an alternate configuration directory")
	initialise = flag.Bool("init", false, "Run first time setup")
)

func main() {
	flag.Parse()

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
