package model

import (
	"testing"

	"time"

	"github.com/stretchr/testify/assert"
)

func TestFindInvite(t *testing.T) {
	db := pg()

	user := NewUser()
	device := NewDevice("one", "1", user.ID)
	invite := NewInvite(NewULID(), device.ID, time.Hour)

	assert.Nil(t, user.Save(db))
	assert.Nil(t, device.Save(db))
	assert.Nil(t, invite.Save(db))

	invite2, err := FindInvite(db, invite.ID)
	assert.Nil(t, err)

	assert.Equal(t, invite.ID, invite2.ID)
	assert.WithinDuration(t, invite.Expires, invite2.Expires, time.Second)
}
