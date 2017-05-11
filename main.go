package main

import (
	"io"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"log"
)

func register(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "moo\n")
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/foo", register).
		Methods("GET")

	srv := &http.Server{
		Addr:         "localhost:8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      router,
	}

	log.Fatal(srv.ListenAndServe())
}
