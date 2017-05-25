package model

import (
	"crypto/rand"

	"github.com/oklog/ulid"
)

// NewULID generates a new ULID - a lexically sortable UUID
func NewULID() string {
	return ulid.MustNew(ulid.Now(), rand.Reader).String()
}
