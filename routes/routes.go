package routes

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

var db *sql.DB

// Handler for all the application routes
func Handler(_db *sql.DB) http.HandlerFunc {
	router := mux.NewRouter()

	db = _db

	router.HandleFunc("/chain", AddLink).
		Methods(http.MethodPost)

	router.Handle("/invite/{key}", NewInviteHandler()).
		Methods(http.MethodGet)

	router.HandleFunc("/about", About).
		Methods(http.MethodGet)

	router.HandleFunc("/", CORS).
		Methods(http.MethodOptions)

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		router.ServeHTTP(w, r)
	}
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
