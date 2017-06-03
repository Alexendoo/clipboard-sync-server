package routes

import (
	"crypto/sha256"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"bytes"

	"github.com/Alexendoo/sync/model"
	"golang.org/x/crypto/ed25519"
)

type genericLink struct {
	Type  string          `json:"type"`
	Body  json.RawMessage `json:"body"`
	SeqNo uint            `json:"seqno"`
	Prev  []byte          `json:"prev"`
}

type newKeyBody struct {
	PubKey ed25519.PublicKey `json:"pkey"`
}

// AddLink adds a link to a users signature chain
func AddLink(w http.ResponseWriter, userid string, body, signature, pubkey []byte) bool {
	var link genericLink
	err := json.Unmarshal(body, &link)
	if err != nil {
		badRequest(w)
		return false
	}

	valid := len(pubkey) == ed25519.PublicKeySize &&
		ed25519.Verify(pubkey, body, signature)
	if !valid {
		forbidden(w)
		log.Printf("invalid signature\n")
		return false
	}

	verified := false
	switch link.Type {
	case "root":
		verified = verifyRoot(w, &link, pubkey)
	case "new_device":
		verified = verifyNewDevice(w, &link, pubkey, userid)
	default:
		badRequest(w)
		verified = false
	}
	if !verified {
		return false
	}

	err = model.NewLink(body, signature, pubkey, userid, link.SeqNo).Save(db)
	if err != nil {
		serverError(w)
		return false
	}

	return true
}

func verifyRoot(w http.ResponseWriter, link *genericLink, pubkey ed25519.PublicKey) bool {
	if link.SeqNo != 0 {
		badRequest(w)
		return false
	}

	var newKey newKeyBody
	err := json.Unmarshal(link.Body, &newKey)
	if err != nil {
		return false
	}

	return bytes.Equal(pubkey, newKey.PubKey)
}

func verifyNewDevice(w http.ResponseWriter, link *genericLink, pubkey ed25519.PublicKey, userid string) bool {
	device, err := model.FindDevice(db, pubkey)
	if err != nil {
		if err == sql.ErrNoRows {
			notFound(w)
		} else {
			serverError(w)
		}
		return false
	}

	if device.UserID != userid {
		forbidden(w)
		return false
	}

	lastLink, err := model.LastLink(db, userid)
	if err != nil {
		serverError(w)
		return false
	}

	if link.SeqNo != lastLink.SeqNo+1 {
		httpError(w, http.StatusConflict)
		return false
	}

	checksum := sha256.Sum256(lastLink.Body)
	if !bytes.Equal(link.Prev, checksum[:]) {
		badRequest(w)
		return false
	}

	return true
}
