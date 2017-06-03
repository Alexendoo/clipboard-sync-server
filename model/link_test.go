package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/ed25519"
)

func TestLastLink(t *testing.T) {
	db := pg()
	user := NewUser()
	assert.Nil(t, user.Save(db))

	pub, _, _ := ed25519.GenerateKey(nil)
	device := NewDevice(pub, "device", "", user.ID)
	assert.Nil(t, device.Save(db))

	link0 := NewLink([]byte("{}"), []byte(""), pub, user.ID, 0)
	assert.Nil(t, link0.Save(db))

	_link, err := LastLink(db, user.ID)
	assert.Equal(t, link0, _link)
	assert.Nil(t, err)

	link1 := NewLink([]byte("{}"), []byte(""), pub, user.ID, 1)
	assert.Nil(t, link1.Save(db))

	_link, err = LastLink(db, user.ID)
	assert.Equal(t, link1, _link)
	assert.Nil(t, err)
}
