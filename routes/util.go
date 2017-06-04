package routes

import (
	"encoding/base64"
	"net/http"
)

func empty(strs ...string) bool {
	for _, str := range strs {
		if str == "" {
			return true
		}
	}
	return false
}

func decodeHeader(r *http.Request, header string) []byte {
	encoded := r.Header.Get(header)
	if len(encoded) == 0 {
		return nil
	}

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return nil
	}

	return decoded
}
