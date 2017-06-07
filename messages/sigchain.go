package messages

import (
	"golang.org/x/crypto/ed25519"
)

//go:generate protoc --go_out=. *.proto

// Verify the signature of the message is valid for the given public key
func (s *Signed) Verify() bool {
	if len(s.PublicKey) != ed25519.PublicKeySize {
		return false
	}

	return ed25519.Verify(s.PublicKey, s.Body, s.Signature)
}
