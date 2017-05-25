package model

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindDevice(t *testing.T) {
	db := pg()

	u := NewUser()
	assert.Nil(t, u.Save(db))

	d := NewDevice("dev", "token", u.ID)
	assert.Nil(t, d.Save(db))

	d2, err := FindDevice(db, d.ID)
	assert.Nil(t, err)

	assert.Equal(t, d, d2)
}

func TestDevice_Save(t *testing.T) {
	db := pg()

	u := NewUser()
	assert.Nil(t, u.Save(db))

	d := NewDevice("foo", "token", u.ID)
	assert.Nil(t, d.Save(db))
}

func TestDevice_Delete(t *testing.T) {
	db := pg()

	u := NewUser()
	assert.Nil(t, u.Save(db))

	d := NewDevice("bar", "t", u.ID)
	assert.Nil(t, d.Save(db))

	assert.Nil(t, d.Delete(db))
	_, err := FindDevice(db, d.ID)
	assert.Equal(t, err, sql.ErrNoRows)
}
