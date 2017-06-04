package routes

import (
	"database/sql"
	"log"

	"github.com/Alexendoo/clipboard-sync-server/testutil"
)

func init() {
	var err error
	db, err = sql.Open("postgres", testutil.URL())
	if err != nil {
		log.Fatal(err)
	}
}
