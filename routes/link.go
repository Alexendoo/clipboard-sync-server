package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"io/ioutil"

	"database/sql"

	"github.com/Alexendoo/sync/model"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/ed25519"
)

type genericLink struct {
	Type  string          `json:"type"`
	Body  json.RawMessage `json:"body"`
	SeqNo uint            `json:"seqno"`
	Prev  string          `json:"prev"`
}

type newKey struct {
	PubKey    ed25519.PublicKey `json:"pkey"`
	UserIDSig []byte            `json:"uid_sig"`
}

// AddLink adds a link to a users signature chain
func AddLink(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["uid"]

	log.Printf("uid: %#+v\n", uid)

	reader := http.MaxBytesReader(w, r.Body, 10)
	body, err := ioutil.ReadAll(reader)
	if err != nil {
		badRequest(w)
		return
	}

	link := &genericLink{}
	err = json.Unmarshal(body, link)
	if err != nil {
		badRequest(w)
		return
	}

	signature := decodeHeader(r, "Sync-Sig")
	pubkey := decodeHeader(r, "Sync-PKey")
	if len(signature) != ed25519.SignatureSize || len(pubkey) != ed25519.PublicKeySize {
		badRequest(w)
		return
	}

	valid := ed25519.Verify(pubkey, body, signature)
	if !valid {
		forbidden(w)
		return
	}

	user, err := model.FindDevice(db, pubkey)
	if err != nil {
		if err == sql.ErrNoRows {
			notFound(w)
		} else {
			serverError(w)
		}
		return
	}

	if user.UserID != uid {
		forbidden(w)
		return
	}

	switch link.Type {
	default:
		badRequest(w)
		return
	}
}
