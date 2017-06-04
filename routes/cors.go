package routes

import (
	"net/http"
)

// CORS - Cross Origin Resource Requests, enables browsers to use the APIs exposed
// by Sync: https://developer.mozilla.org/en-US/docs/Web/HTTP/Access_control_CORS
func CORS(w http.ResponseWriter, r *http.Request) {
	h := w.Header()
	h.Set("Access-Control-Allow-Methods", "GET, POST")
	h.Set("Access-Control-Allow-Headers", "Content-Type")
}
