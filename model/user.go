package model

import (
	"database/sql"
)

type User struct {
	ID string
}

func NewUser() *User {
	return &User{
		ID: NewULID(),
	}
}

func FindUser(db *sql.DB, id string) (*User, error) {
	row := db.QueryRow("SELECT * FROM users WHERE id = $1", id)
	var ID string
	err := row.Scan(&ID)
	return &User{ID}, err
}

func (u *User) Save(db *sql.DB) error {
	_, err := db.Exec(
		"INSERT INTO users VALUES ($1)",
		u.ID,
	)

	return err
}
