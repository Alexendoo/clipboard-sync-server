package model

import (
	"testing"
)

func TestFindInvite(t *testing.T) {
	db := pg()

	user := NewUser()
	device := NewDevice("one", "1", user.ID)
	invite := NewInvite(NewULID(), device.ID)

	user.Save(db)
	device.Save(db)
	invite.Save(db)

	invite2, err := FindInvite(db, invite.ID)
	if err != nil {
		t.Errorf("FindInvite() error = %v", err)
	}

	if !invite.Expires.Equal(invite2.Expires) {
		t.Errorf("unequal expires: %v, %v", invite.Expires, invite2.Expires)
	}
}
