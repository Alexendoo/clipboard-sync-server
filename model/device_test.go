package model

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/ed25519"
)

func TestFindDevice(t *testing.T) {
	db := pg()

	u := NewUser()
	assert.Nil(t, u.Save(db))

	pub, _, err := ed25519.GenerateKey(nil)
	assert.Nil(t, err)

	d := NewDevice(pub, "dev", "token", u.ID)
	assert.Nil(t, d.Save(db))

	d2, err := FindDevice(db, d.PKey)
	assert.Nil(t, err)

	assert.Equal(t, d, d2)
}

func TestDevice_Save(t *testing.T) {
	db := pg()

	u := NewUser()
	assert.Nil(t, u.Save(db))

	pub, _, err := ed25519.GenerateKey(nil)
	assert.Nil(t, err)

	d := NewDevice(pub, "foo", "token", u.ID)
	assert.Nil(t, d.Save(db))
}

func TestDevice_Delete(t *testing.T) {
	db := pg()

	u := NewUser()
	assert.Nil(t, u.Save(db))

	pub, _, err := ed25519.GenerateKey(nil)
	assert.Nil(t, err)

	d := NewDevice(pub, "bar", "t", u.ID)
	assert.Nil(t, d.Save(db))

	assert.Nil(t, d.Delete(db))
	_, err = FindDevice(db, d.PKey)
	assert.Equal(t, err, sql.ErrNoRows)
}
