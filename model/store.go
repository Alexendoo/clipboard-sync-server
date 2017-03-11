package model

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Store struct {
	db *sql.DB
}

type UserStore interface {
	AddUser(user *User) error
}

func OpenStore(path string) (store *Store, err error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return
	}

	rows, err := db.Query(`
SELECT name FROM sqlite_master
WHERE type='table' AND name='config'
`)
	dbExists := rows.Next()
	rows.Close()

	if dbExists {
		return &Store{db}, nil
	}

	tx, err := db.Begin()
	if err != nil {
		return
	}
	defer tx.Rollback()

	_, err = tx.Exec(`
CREATE TABLE IF NOT EXISTS config (
	id INTEGER PRIMARY KEY CHECK (id = 0) DEFAULT 0,
	version INTEGER DEFAULT 0
);

INSERT INTO config DEFAULT VALUES;

CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY,

	name STRING COLLATE NOCASE,

	pass_version INTEGER,
	pass_salt BLOB,
	pass_key BLOB
);

CREATE UNIQUE INDEX IF NOT EXISTS ux_users_name
ON users (name COLLATE NOCASE);
`)
	if err != nil {
		return
	}

	return &Store{db}, tx.Commit()
}

func (s *Store) AddUser(user *User) error {

	return nil
}
