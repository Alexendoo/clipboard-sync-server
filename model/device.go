package model

import (
	"database/sql"
)

// Device is a data model of a users device (browser extension, mobile app)
type Device struct {
	ID       string
	Name     string
	FCMToken string
	UserID   string
}

// NewDevice returns a new instance of Device
func NewDevice(name, token, userID string) *Device {
	return &Device{
		ID:       NewULID(),
		Name:     name,
		FCMToken: token,
		UserID:   userID,
	}
}

// FindDevice finds the Device from the database with the provided ID
func FindDevice(db *sql.DB, id string) (*Device, error) {
	row := db.QueryRow("SELECT * FROM devices WHERE id = $1", id)
	var (
		ID       string
		Name     string
		FCMToken string
		UserID   string
	)
	err := row.Scan(&ID, &Name, &FCMToken, &UserID)

	return &Device{ID, Name, FCMToken, UserID}, err
}

// Save the Device into the database
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

// Delete the Device entry from the database
func (d *Device) Delete(db *sql.DB) error {
	_, err := db.Exec(
		"DELETE FROM devices WHERE id = $1",
		d.ID,
	)

	return err
}
