package routes

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/ed25519"
)

func TestGetPayload(t *testing.T) {
	pub, priv, err := ed25519.GenerateKey(nil)
	assert.Nil(t, err)

	payload := make([]byte, 32)
	_, err = io.ReadFull(rand.Reader, payload)
	assert.Nil(t, err)

	sig := ed25519.Sign(priv, payload)

	buf := &bytes.Buffer{}
	buf.WriteByte(0)
	buf.Write(pub)
	buf.Write(sig)
	buf.Write(payload)

	valid := buf.Bytes()

	t.Run("Valid", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(valid))
		recorder := httptest.NewRecorder()

		_pub, _payload := getPayload(recorder, request)
		assert.Equal(t, payload, _payload)
		assert.Equal(t, pub, _pub)
		assert.Equal(t, recorder.Code, http.StatusOK)
	})

	for i := range valid {
		t.Run(fmt.Sprintf("Invalid %d", i), func(t *testing.T) {
			invalid := make([]byte, len(valid))
			copy(invalid, valid)
			invalid[i] ^= 1

			request := httptest.NewRequest(
				http.MethodPost,
				"/",
				bytes.NewReader(invalid),
			)
			recorder := httptest.NewRecorder()

			_pub, _payload := getPayload(recorder, request)
			assert.Nil(t, _pub)
			assert.Nil(t, _payload)
			if i == 0 {
				assert.Equal(t, recorder.Code, http.StatusBadRequest)
			} else {
				assert.Equal(t, recorder.Code, http.StatusForbidden)
			}
		})
	}
}

func TestGetPayloadBadSize(t *testing.T) {
	nums := []int64{0, 1, 95, 96, 97, 98, 8192, 8193, 65536}
	for _, num := range nums {
		t.Run(fmt.Sprintf("Size %d", num), func(t *testing.T) {
			limit := io.LimitReader(rand.Reader, num-1)
			reader := io.MultiReader(bytes.NewReader([]byte{0}), limit)

			request := httptest.NewRequest(http.MethodPost, "/", reader)
			recorder := httptest.NewRecorder()

			pub, payload := getPayload(recorder, request)
			assert.Nil(t, pub)
			assert.Nil(t, payload)
			assert.Equal(t, http.StatusBadRequest, recorder.Code)
		})
	}
}
