package routes

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Alexendoo/clipboard-sync-server/messages"
	"github.com/Alexendoo/clipboard-sync-server/model"
	"github.com/golang/protobuf/proto"
	"golang.org/x/crypto/ed25519"
)

type linkRequest struct {
	Link      []byte            `json:"link"`
	Signature []byte            `json:"sig"`
	PubKey    ed25519.PublicKey `json:"pkey"`
}

type genericLink struct {
	Type  string          `json:"type"`
	Body  json.RawMessage `json:"body"`
	SeqNo uint            `json:"seqno"`
	Prev  []byte          `json:"prev"`
}

type newDeviceBody struct {
	Name     string            `json:"name"`
	PubKey   ed25519.PublicKey `json:"pkey"`
	FCMToken string            `json:"fcm"`
}

// AddLink adds a link to a users signature chain
func AddLink(w http.ResponseWriter, r *http.Request) {
	reader := http.MaxBytesReader(w, r.Body, 10000)
	body, err := ioutil.ReadAll(reader)
	if err != nil {
		badRequest(w)
		return
	}

	signed := &messages.Signed{}
	err = proto.Unmarshal(body, signed)
	if err != nil {
		badRequest(w)
		return
	}

	if !signed.Verify() {
		forbidden(w)
		return
	}

	link := &messages.Link{}
	err = proto.Unmarshal(signed.Body, link)
	if err != nil {
		badRequest(w)
		return
	}

	switch link.Body.(type) {
	case *messages.Link_NewDevice:
		addNewDevice(w, signed, link)
	default:
		badRequest(w)
	}
}

func addNewDevice(w http.ResponseWriter, signed *messages.Signed, link *messages.Link) {
	if link.SequenceNumber == 0 {
		addRootDevice(w, signed, link)
		return
	}

	signatory, err := model.FindDevice(db, signed.PublicKey)
	if err != nil {
		badRequest(w)
		return
	}

	lastLink, err := model.LastLink(db, signatory.UserID)
	if err != nil {
		serverError(w)
		return
	}

	if link.SequenceNumber != lastLink.SeqNo+1 {
		httpError(w, http.StatusConflict)
		return
	}

	checksum := sha256.Sum256(lastLink.Body)
	if !bytes.Equal(link.Prev, checksum[:]) {
		badRequest(w)
		return
	}

	newDevice := link.Body.(*messages.Link_NewDevice).NewDevice

	device := model.NewDevice(newDevice.PublicKey, newDevice.FCMToken, signatory.UserID)
	err = device.Save(db)
	if err != nil {
		serverError(w)
		return
	}

	savedLink := model.NewLink(signed.Body, signed.Signature, signed.PublicKey, signatory.UserID, link.SequenceNumber)
	err = savedLink.Save(db)
	if err != nil {
		serverError(w)
		device.Delete(db)
		return
	}
}

func addRootDevice(w http.ResponseWriter, signed *messages.Signed, link *messages.Link) {
	newDevice := link.Body.(*messages.Link_NewDevice).NewDevice

	if !bytes.Equal(newDevice.PublicKey, signed.PublicKey) {
		badRequest(w)
		return
	}

	user := model.NewUser()
	err := user.Save(db)
	if err != nil {
		serverError(w)
		return
	}

	device := model.NewDevice(newDevice.PublicKey, newDevice.FCMToken, user.ID)
	err = device.Save(db)
	if err != nil {
		serverError(w)
		user.Delete(db)
		return
	}

	savedLink := model.NewLink(
		signed.Body,
		signed.Signature,
		signed.PublicKey,
		user.ID,
		link.SequenceNumber,
	)
	err = savedLink.Save(db)
	if err != nil {
		serverError(w)
		device.Delete(db)
		user.Delete(db)
		return
	}
}
