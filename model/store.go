package model

import (
	"bytes"

	"encoding/gob"

	"github.com/boltdb/bolt"
)

type Store struct {
	db         bolt.DB
	userBucket []byte
}

type UserStore interface {
	AddUser(user *User) error
}

func (s *Store) AddUser(user *User) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(s.userBucket)

		var buffer bytes.Buffer
		encoder := gob.NewEncoder(&buffer)
		encoder.Encode(user)

		err := b.Put([]byte(user.ID), buffer.Bytes())
		return err
	})
}
