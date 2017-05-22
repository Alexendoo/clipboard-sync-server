package model

import (
	"database/sql"
	"reflect"
	"testing"
)

func TestFindUser(t *testing.T) {
	db := pg()
	u := NewUser()
	u.Save(db)

	u2, err := FindUser(db, u.ID)
	if err != nil {
		t.Errorf("FindUser() error = %v", err)
	}
	if !reflect.DeepEqual(u, u2) {
		t.Error("u != u2")
	}
}

func TestUser_Save(t *testing.T) {
	db := pg()
	u := NewUser()
	err := u.Save(db)
	if err != nil {
		t.Errorf("User.Save() error = %v", err)
	}
}

func TestUser_Delete(t *testing.T) {
	db := pg()
	u := NewUser()
	u.Save(db)

	err := u.Delete(db)
	if err != nil {
		t.Errorf("User.Delete() error = %v", err)
	}

	_, err = FindUser(db, u.ID)
	if err != sql.ErrNoRows {
		t.Errorf("FindUser() error = %v", err)
	}
}
