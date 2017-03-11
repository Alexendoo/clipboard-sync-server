package model

type User struct {
	Name string

	PassVersion int
	PassSalt    []byte
	PassKey     []byte
}
