package model

import kallax "gopkg.in/src-d/go-kallax.v1"

type Device struct {
	kallax.Model `table:"devices"`
	ID           kallax.ULID `pk:""`

	Name string

	User *User `fk:"user_id,inverse" json:"-"`
}

func newDevice(name string, user *User) *Device {
	return &Device{
		ID:   kallax.NewULID(),
		Name: name,
		User: user,
	}
}
