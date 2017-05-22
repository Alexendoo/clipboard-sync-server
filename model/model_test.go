package model

import (
	"database/sql"
	"log"
)

func pg() *sql.DB {
	db, err := sql.Open("postgres", "postgres://devel:devel@localhost/sync?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
