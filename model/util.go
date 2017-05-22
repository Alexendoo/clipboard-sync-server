package model

import (
	"math/rand"
	"sync"
	"time"

	"github.com/oklog/ulid"
)

var randPool = &sync.Pool{
	New: func() interface{} {
		source := rand.NewSource(time.Now().UnixNano())
		return rand.New(source)
	},
}

// NewULID generates a new ULID - a lexically sortable UUID
func NewULID() string {
	entropy := randPool.Get().(*rand.Rand)
	return ulid.MustNew(ulid.Now(), entropy).String()
}
