package model

import (
	"log"
	"testing"
)

func TestNewUser(t *testing.T) {
	db := openDb()

	userStore := NewUserStore(db)
	u := NewUser()

	log.Println(userStore.Insert(u))

	uq := NewUserQuery().FindByID(u.ID)
	res, _ := userStore.FindOne(uq)

	log.Printf("res: %#+v\n", res)

	deviceStore := NewDeviceStore(db)
	d := NewDevice("moo", "chrome", u)

	log.Println(deviceStore.Insert(d))
}
