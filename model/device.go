package model

import (
	"database/sql"

	"golang.org/x/crypto/ed25519"
)

// Device is a data model of a users device (browser extension, mobile app)
type Device struct {
	PKey     ed25519.PublicKey
	FCMToken string
	UserID   string
}

// NewDevice returns a new instance of Device
func NewDevice(pubkey ed25519.PublicKey, token, userID string) *Device {
	return &Device{
		PKey:     pubkey,
		FCMToken: token,
		UserID:   userID,
	}
}

// FindDevice finds the Device from the database with the provided ID
func FindDevice(db *sql.DB, pkey ed25519.PublicKey) (*Device, error) {
	row := db.QueryRow("SELECT * FROM devices WHERE public_key = $1", pkey)
	var (
		PKey     ed25519.PublicKey
		FCMToken string
		UserID   string
	)
	err := row.Scan(&PKey, &FCMToken, &UserID)

	return &Device{PKey, FCMToken, UserID}, err
}

// Save the Device into the database
func (d *Device) Save(db *sql.DB) error {
	_, err := db.Exec(
		"INSERT INTO devices VALUES ($1, $2, $3)",
		d.PKey,
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
