package model

import (
	"reflect"
	"testing"
)

func TestFindDevice(t *testing.T) {
	db := pg()

	u := NewUser()
	u.Save(db)

	d := NewDevice("dev", "token", u.ID)
	d.Save(db)

	d2, err := FindDevice(db, d.ID)
	if err != nil {
		t.Errorf("FindDevice() error = %v", err)
	}
	if !reflect.DeepEqual(d, d2) {
		t.Fail()
	}
}

func TestDevice_Save(t *testing.T) {
	db := pg()

	u := NewUser()
	u.Save(db)

	d := NewDevice("foo", "token", u.ID)
	err := d.Save(db)
	if err != nil {
		t.Errorf("Device.Save() error = %v", err)
	}
}
