package routes

import (
	"bytes"
	"crypto/sha256"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Alexendoo/clipboard-sync-server/model"
	"golang.org/x/crypto/ed25519"
)

type linkRequest struct {
	Link      json.RawMessage
	Signature []byte
	PubKey    ed25519.PublicKey
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
		// TODO: check resp badRequest(w)
		return
	}

	var req linkRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		badRequest(w)
		return
	}

	valid := len(req.PubKey) == ed25519.PublicKeySize &&
		ed25519.Verify(req.PubKey, req.Link, req.Signature)
	if !valid {
		forbidden(w)
		log.Printf("invalid signature\n")
		return
	}

	var link genericLink
	err = json.Unmarshal(req.Link, &link)
	if err != nil {
		badRequest(w)
		return
	}

	switch link.Type {
	case "root":
		addRoot(w, &link, &req)
	case "new_device":
		addNewDevice(w, &link, &req)
	default:
		badRequest(w)
	}
}

func addRoot(w http.ResponseWriter, link *genericLink, req *linkRequest) {
	if link.SeqNo != 0 {
		badRequest(w)
		return
	}

	var newDevice newDeviceBody
	err := json.Unmarshal(link.Body, &newDevice)
	if err != nil {
		badRequest(w)
		return
	}

	user := model.NewUser()
	err = user.Save(db)
	if err != nil {
		serverError(w)
		return
	}

	device := model.NewDevice(req.PubKey, newDevice.FCMToken, user.ID)
	err = device.Save(db)
	if err != nil {
		serverError(w)
		user.Delete(db)
		return
	}

	savedLink := model.NewLink(req.Link, req.Signature, req.PubKey, user.ID, link.SeqNo)
	err = savedLink.Save(db)
	if err != nil {
		serverError(w)
		user.Delete(db)
		device.Delete(db)
		return
	}

	resp, _ := json.Marshal(&savedLink)
	w.Write(resp)
}

func addNewDevice(w http.ResponseWriter, link *genericLink, req *linkRequest) {
	var newDevice newDeviceBody
	err := json.Unmarshal(link.Body, &newDevice)
	if err != nil {
		badRequest(w)
		return
	}

	signatory, err := model.FindDevice(db, req.PubKey)
	if err != nil {
		if err == sql.ErrNoRows {
			badRequest(w)
		} else {
			serverError(w)
		}
		return
	}

	lastLink, err := model.LastLink(db, signatory.UserID)
	if err != nil {
		serverError(w)
		return
	}

	if link.SeqNo != lastLink.SeqNo+1 {
		httpError(w, http.StatusConflict)
		return
	}

	checksum := sha256.Sum256(lastLink.Body)
	if !bytes.Equal(link.Prev, checksum[:]) {
		badRequest(w)
		return
	}

	device := model.NewDevice(newDevice.PubKey, newDevice.FCMToken, signatory.UserID)
	err = device.Save(db)
	if err != nil {
		serverError(w)
		return
	}

	savedLink := model.NewLink(req.Link, req.Signature, req.PubKey, signatory.UserID, link.SeqNo)
	err = savedLink.Save(db)
	if err != nil {
		serverError(w)
		device.Delete(db)
		return
	}
}
