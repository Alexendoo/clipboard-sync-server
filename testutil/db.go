package testutil

import (
	"os"
)

// URL returns the postgres database URL
func URL() string {
	url, ok := os.LookupEnv("POSTGRES_URL")
	if !ok {
		url = "postgres://devel:devel@localhost/sync?sslmode=disable"
	}

	return url
}
