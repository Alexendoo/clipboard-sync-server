package model

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Store struct {
	db *sql.DB
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
		return &Store{db}, err
	}

	tx, err := db.Begin()
	if err != nil {
		return
	}
	defer tx.Rollback()

	_, err = tx.Exec(`
PRAGMA foreign_keys = ON;

CREATE TABLE config (
	version INTEGER PRIMARY KEY
);

INSERT INTO config (version)
VALUES (0);


CREATE TABLE users (
	id INTEGER PRIMARY KEY,

	username STRING COLLATE NOCASE,
	pass_key BLOB,
	pass_salt BLOB,
	pass_version INTEGER
);

CREATE UNIQUE INDEX ux_users_username
ON users (username COLLATE NOCASE);


CREATE TABLE devices (
	id INTEGER PRIMARY KEY,
	auth_key BLOB NOT NULL,
	userid INTEGER NOT NULL,

	FOREIGN KEY (userid) REFERENCES users (id)
);

CREATE INDEX ix_devices_userid
ON devices (userid);
`)
	if err != nil {
		return
	}

	return &Store{db}, tx.Commit()
}
