package routes

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

var db *sql.DB

// Handler for all the application routes
func Handler(_db *sql.DB) http.Handler {
	router := mux.NewRouter()

	db = _db

	n := negroni.New()

	router.HandleFunc("/register", Register).
		Methods(http.MethodPost)

	router.HandleFunc("/link/{uid}", AddLink).
		Methods(http.MethodPost)

	router.HandleFunc("/invite/{device:src|dest}/{key}", InviteGet).
		Methods(http.MethodGet)
	router.HandleFunc("/invite/{device:src|dest}/{key}", InvitePost).
		Methods(http.MethodPost)

	n.UseFunc(jsonHeader)
	n.UseHandler(router)

	return n
}

func jsonHeader(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	w.Header().Set("Content-Type", "application/json")
	next(w, r)
}

type httpErr struct {
	Error string
}

func httpError(w http.ResponseWriter, code int) {
	err := &httpErr{http.StatusText(code)}
	bytes, _ := json.Marshal(err)
	w.WriteHeader(code)
	w.Write(bytes)
}

func badRequest(w http.ResponseWriter) {
	httpError(w, http.StatusBadRequest)
}

func forbidden(w http.ResponseWriter) {
	httpError(w, http.StatusForbidden)
}

func serverError(w http.ResponseWriter) {
	httpError(w, http.StatusInternalServerError)
}

func notFound(w http.ResponseWriter) {
	httpError(w, http.StatusNotFound)
}
