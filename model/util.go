package model

import (
	"github.com/oklog/ulid"
	"sync"
	"math/rand"
	"time"
)

var randPool = &sync.Pool{
	New: func() interface{} {
		source := rand.NewSource(time.Now().UnixNano())
		return rand.New(source)
	},
}

// NewULID generates a new ULID - a lexically sortable UUID
func NewULID() ulid.ULID {
	entropy := randPool.Get().(rand.Rand)
	return ulid.MustNew(ulid.Now(), entropy)
}
