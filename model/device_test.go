package model

import (
	"testing"
)

func TestDevice_Save(t *testing.T) {
	db := getDB()

	u := NewUser()
	err := u.Save(db)
	if err != nil {
		t.Errorf("User.Save() error = %v", err)
	}
	d := NewDevice("foo", "token", u.ID)
	err = d.Save(db)
	if err != nil {
		t.Errorf("Device.Save() error = %v", err)
	}
}
