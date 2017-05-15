package model

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func openDb() *sql.DB {
	db, err := sql.Open("postgres", "postgres://devel:devel@localhost/sync?sslmode=disable")
	if err != nil {
		panic(err)
	}
	return db
}
