package routes

import (
	"database/sql"
	"log"
)

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://devel:devel@localhost/sync?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
}
