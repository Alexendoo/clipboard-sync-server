package model

import (
	"github.com/oklog/ulid"
	"database/sql"
)

type Device struct {
	ID ulid.ULID

	Name     string
	FCMToken string

	UserID ulid.ULID
}

func NewDevice(name, token string, userID ulid.ULID) *Device {
	return &Device{
		ID:       NewULID(),
		Name:     name,
		FCMToken: token,
		UserID:   userID,
	}
}

func (d *Device) Save(db *sql.DB) {

}
