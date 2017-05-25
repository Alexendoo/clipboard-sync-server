package model

import "time"
import "database/sql"

// Invite is a request to add another Device to a User
type Invite struct {
	ID       string
	Expires  time.Time
	DeviceID string
}

// NewInvite creates an invite valid for one hour
func NewInvite(id, deviceID string) *Invite {
	return &Invite{
		ID:       id,
		Expires:  time.Now().Add(1 * time.Hour),
		DeviceID: deviceID,
	}
}

// FindInvite finds the Invite from the database with the provided ID
func FindInvite(db *sql.DB, id string) (*Invite, error) {
	row := db.QueryRow("SELECT * FROM invites WHERE id = $1", id)
	var (
		ID       string
		Expires  time.Time
		DeviceID string
	)
	err := row.Scan(&ID, &Expires, &DeviceID)

	return &Invite{ID, Expires, DeviceID}, err
}

// Save the Invite into the database
func (i *Invite) Save(db *sql.DB) error {
	_, err := db.Exec(
		"INSERT INTO invites VALUES ($1, $2, $3)",
		i.ID,
		i.Expires,
		i.DeviceID,
	)

	return err
}
