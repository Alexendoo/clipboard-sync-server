package routes

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/ed25519"
)

func TestRegisterUser(t *testing.T) {
	assert := assert.New(t)

	pub, priv, err := ed25519.GenerateKey(nil)
	assert.NoError(err)

	linkBody, err := json.Marshal(&newKeyBody{
		PubKey: pub,
	})
	assert.NoError(err)

	link := genericLink{
		Type:  "root",
		Body:  linkBody,
		SeqNo: 0,
		Prev:  nil,
	}
	linkBytes, err := json.Marshal(&link)
	assert.NoError(err)

	sig := ed25519.Sign(priv, linkBytes)

	req := newUserRequest{
		Link:      linkBytes,
		FCMToken:  "1234",
		Signature: sig,
		PubKey:    pub,
	}
	body, err := json.Marshal(&req)
	assert.NoError(err)

	request := httptest.NewRequest(http.MethodPost, "/register", bytes.NewReader(body))
	recorder := httptest.NewRecorder()

	hand := Handler(db)
	hand.ServeHTTP(recorder, request)

	assert.Equal(recorder.Code, http.StatusOK, recorder.Body.String())
}
