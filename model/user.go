package model

import kallax "gopkg.in/src-d/go-kallax.v1"

type User struct {
	kallax.Model `table:"users"`
	ID           kallax.ULID `pk:""`
}

func newUser() *User {
	return &User{
		ID: kallax.NewULID(),
	}
}
