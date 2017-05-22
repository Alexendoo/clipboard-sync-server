package model

import (
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
		t.Fail()
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
