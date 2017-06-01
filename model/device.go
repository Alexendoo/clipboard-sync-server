package model

import (
	"database/sql"

	"golang.org/x/crypto/ed25519"
)

// Device is a data model of a users device (browser extension, mobile app)
type Device struct {
	PKey     ed25519.PublicKey
	Name     string
	FCMToken string
	UserID   string
}

// NewDevice returns a new instance of Device
func NewDevice(pubkey ed25519.PublicKey, name, token, userID string) *Device {
	return &Device{
		PKey:     pubkey,
		Name:     name,
		FCMToken: token,
		UserID:   userID,
	}
}

// FindDevice finds the Device from the database with the provided ID
func FindDevice(db *sql.DB, pkey ed25519.PublicKey) (*Device, error) {
	row := db.QueryRow("SELECT * FROM devices WHERE public_key = $1", pkey)
	var (
		PKey     ed25519.PublicKey
		Name     string
		FCMToken string
		UserID   string
	)
	err := row.Scan(&PKey, &Name, &FCMToken, &UserID)

	return &Device{PKey, Name, FCMToken, UserID}, err
}

// Save the Device into the database
func (d *Device) Save(db *sql.DB) error {
	_, err := db.Exec(
		"INSERT INTO devices VALUES ($1, $2, $3, $4)",
		d.PKey,
		d.Name,
		d.FCMToken,
		d.UserID,
	)
	return err
}

// Delete the Device entry from the database
func (d *Device) Delete(db *sql.DB) error {
	_, err := db.Exec(
		"DELETE FROM devices WHERE public_key = $1",
		d.PKey,
	)

	return err
}
