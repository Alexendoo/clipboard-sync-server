package model

import kallax "gopkg.in/src-d/go-kallax.v1"

type Sigchain struct {
}

type Link struct {
	kallax.Model `table:"sigchain"`
	ID           kallax.ULID `pk:""`

	Data           []byte `kallax:"link"`
	SequenceNumber int

	User *User `fk:"user_id,inverse"`
}
