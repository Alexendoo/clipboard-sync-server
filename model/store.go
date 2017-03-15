package model

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Store struct {
	db    *sql.DB
	stmts *stmts
}

type stmts struct {
	addUser *sql.Stmt
	getUser *sql.Stmt
}

func OpenStore(path string) (*Store, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(`
SELECT name FROM sqlite_master
WHERE type='table' AND name='config'
`)
	dbExists := rows.Next()
	rows.Close()

	if !dbExists {
		initialiseDB(db)
	}

	stmts, err := prepareStmts(db)

	return &Store{db, stmts}, err
}

func initialiseDB(db *sql.DB) error {
	_, err := db.Exec(`
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
	return err
}

func prepareStmts(db *sql.DB) (statements *stmts, err error) {
	addUser, err := db.Prepare(addUserSQL)
	if err != nil {
		return
	}

	getUser, err := db.Prepare(getUserSQL)
	if err != nil {
		return
	}

	statements = &stmts{
		addUser,
		getUser,
	}

	return
}
