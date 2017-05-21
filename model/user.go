package model

import (
	kallax "gopkg.in/src-d/go-kallax.v1"
	"github.com/oklog/ulid"
)

type User struct {
	kallax.Model `table:"users"`
	ID ulid.ULID `pk:""`
}

func newUser() *User {
	return &User{
		ID: NewULID(),
	}
}
