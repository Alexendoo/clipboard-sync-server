package model

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindUser(t *testing.T) {
	db := pg()
	u := NewUser()
	assert.Nil(t, u.Save(db))

	u2, err := FindUser(db, u.ID)
	assert.Nil(t, err)
	assert.Equal(t, u, u2)
}

func TestUserExists(t *testing.T) {
	db := pg()
	u := NewUser()
	assert.Nil(t, u.Save(db))

	exists, err := UserExists(db, u.ID)
	assert.True(t, exists)
	assert.Nil(t, err)

	exists, err = UserExists(db, "_")
	assert.False(t, exists)
	assert.Nil(t, err)
}

func TestUser_Save(t *testing.T) {
	db := pg()
	u := NewUser()
	assert.Nil(t, u.Save(db))
}

func TestUser_Delete(t *testing.T) {
	db := pg()
	u := NewUser()
	assert.Nil(t, u.Save(db))

	assert.Nil(t, u.Delete(db))

	_, err := FindUser(db, u.ID)
	assert.Equal(t, err, sql.ErrNoRows)
}
