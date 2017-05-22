package model

import (
	"database/sql"
)

type Device struct {
	ID       string
	Name     string
	FCMToken string
	UserID   string
}

func NewDevice(name, token, userID string) *Device {
	return &Device{
		ID:       NewULID(),
		Name:     name,
		FCMToken: token,
		UserID:   userID,
	}
}

func FindDevice(db *sql.DB, id string) (*Device, error) {
	row := db.QueryRow("SELECT * FROM devices WHERE id = $1", id)
	var (
		ID       string
		Name     string
		FCMToken string
		UserID   string
	)
	err := row.Scan(ID, Name, FCMToken, UserID)

	return &Device{ID, Name, FCMToken, UserID}, err
}

func (d *Device) Save(db *sql.DB) error {
	_, err := db.Exec(
		"INSERT INTO devices VALUES ($1, $2, $3, $4)",
		d.ID,
		d.Name,
		d.FCMToken,
		d.UserID,
	)
	return err
}
