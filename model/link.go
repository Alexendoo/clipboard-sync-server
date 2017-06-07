package model

import (
	"database/sql"
)

// Link represents a single entry in a User's sigchain
type Link struct {
	Body []byte

	Signature []byte
	SignedBy  []byte

	UserID string
	SeqNo  uint32
}

// NewLink returns a new Link
func NewLink(Body, Signature, SignedBy []byte, UserID string, SeqNo uint32) *Link {
	return &Link{Body, Signature, SignedBy, UserID, SeqNo}
}

// LastLink returns the last Link in a User's sigchain
func LastLink(db *sql.DB, userid string) (*Link, error) {
	row := db.QueryRow("SELECT * FROM sigchain WHERE user_id = $1 ORDER BY seq_no DESC LIMIT 1", userid)
	var (
		Body      []byte
		Signature []byte
		SignedBy  []byte
		UserID    string
		SeqNo     uint32
	)
	err := row.Scan(
		&Body,
		&Signature,
		&SignedBy,
		&UserID,
		&SeqNo,
	)

	return &Link{Body, Signature, SignedBy, UserID, SeqNo}, err
}

// Save the Link into the database
func (l *Link) Save(db *sql.DB) error {
	_, err := db.Exec(
		"INSERT INTO sigchain VALUES ($1, $2, $3, $4, $5)",
		l.Body,
		l.Signature,
		l.SignedBy,
		l.UserID,
		l.SeqNo,
	)

	return err
}
