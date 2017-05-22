package model

import (
	"database/sql"
	"log"
)

var _pg *sql.DB

func pg() *sql.DB {
	if _pg == nil {
		var err error
		_pg, err = sql.Open("postgres", "postgres://devel:devel@localhost/sync?sslmode=disable")
		if err != nil {
			log.Fatal(err)
		}
	}
	return _pg
}
