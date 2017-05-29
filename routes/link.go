package routes

import (
	"encoding/json"
	"net/http"

	"io/ioutil"

	"io"

	"golang.org/x/crypto/ed25519"
)

type genericLink struct {
	Type  string          `json:"type"`
	Body  json.RawMessage `json:"body"`
	SeqNo uint            `json:"seqno"`
	Prev  string          `json:"prev"`
}

type newKey struct {
	KeyID     ed25519.PublicKey `json:"kid"`
	UserIDSig []byte            `json:"uid_sig"`
}

func getPayload(w http.ResponseWriter, r *http.Request) (key ed25519.PublicKey, payload []byte) {
	const minsize = 2 + ed25519.PublicKeySize + ed25519.SignatureSize
	const maxsize = 8192

	reader := io.LimitReader(r.Body, maxsize)
	body, err := ioutil.ReadAll(reader)

	if err != nil || len(body) <= minsize || len(body) >= maxsize {
		badRequest(w)
		return nil, nil
	}

	version := body[0]
	if version != 0 {
		badRequest(w)
		return nil, nil
	}

	key = body[1:33]
	sig := body[33:97]
	payload = body[97:]

	valid := ed25519.Verify(key, payload, sig)
	if !valid {
		forbidden(w)
		return nil, nil
	}

	return key, payload
}

// AddLink adds a link to a users signature chain, links are formatted as:
//
//   version (1 byte)
//   publicKey (32 bytes)
//   signature (64 bytes)
//   payload (arbitrary bytes)
//
// where version is 0, publicKey is an ed25119 public key, signature
// is ed25519.Sign(publicKey, payload) and payload is a json encoded
// genericLink
func AddLink(w http.ResponseWriter, r *http.Request) {
	_, payload := getPayload(w, r)
	if payload == nil {
		return
	}
}
