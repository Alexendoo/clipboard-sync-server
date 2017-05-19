package model

import (
	"log"
	"testing"
)

func TestNewDevice(t *testing.T) {
	db := openDb()

	userStore := NewUserStore(db).Debug()
	u := NewUser()

	log.Println(userStore.Insert(u))

	uq := NewUserQuery().FindByID(u.ID)
	res, err := userStore.FindOne(uq)

	log.Printf("res: %#+v\n", res)
	log.Printf("err: %#+v\n", err)

	deviceStore := NewDeviceStore(db)
	d := NewDevice("moo", u)

	log.Println(deviceStore.Insert(d))
}
