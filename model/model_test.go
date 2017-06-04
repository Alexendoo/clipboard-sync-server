package model

import (
	"database/sql"
	"log"

	"github.com/Alexendoo/clipboard-sync-server/testutil"
)

var _pg *sql.DB

func pg() *sql.DB {
	if _pg == nil {
		var err error
		_pg, err = sql.Open("postgres", testutil.URL())
		if err != nil {
			log.Fatal(err)
		}
	}
	return _pg
}
