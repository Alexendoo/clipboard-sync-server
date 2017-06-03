package routes

import (
	"io"
	"net/http"
)

// About exposes server information to clients
func About(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `{
		"version": "0.1.0",
		"sender_id": "303334042045"
	}`)
}
