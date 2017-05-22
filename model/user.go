package model

import (
	"database/sql"
)

// User data model
type User struct {
	ID string
}

// NewUser returns a new instance of User
func NewUser() *User {
	return &User{
		ID: NewULID(),
	}
}

// FindUser finds the User from the database with the provided ID
func FindUser(db *sql.DB, id string) (*User, error) {
	row := db.QueryRow("SELECT * FROM users WHERE id = $1", id)
	var ID string
	err := row.Scan(&ID)
	return &User{ID}, err
}

// UserExists is a lightweight check to see if a User with a given ID
// is stored in the database
func UserExists(db *sql.DB, id string) (bool, error) {
	row := db.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)",
		id,
	)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

// Save the User into the database
func (u *User) Save(db *sql.DB) error {
	_, err := db.Exec(
		"INSERT INTO users VALUES ($1)",
		u.ID,
	)
	return err
}

// Delete the User entry from the database
func (u *User) Delete(db *sql.DB) error {
	_, err := db.Exec(
		"DELETE FROM users WHERE id = $1",
		u.ID,
	)
	return err
}
