package routes

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

var db *sql.DB

// Handler for all the application routes
func Handler(_db *sql.DB) http.Handler {
	router := mux.NewRouter()

	db = _db

	router.HandleFunc("/register", Register).
		Methods(http.MethodPost)

	return router
}

func httpError(w http.ResponseWriter, code int) {
	http.Error(w, http.StatusText(code), code)
}

func badRequest(w http.ResponseWriter) {
	httpError(w, http.StatusBadRequest)
}

func serverError(w http.ResponseWriter) {
	httpError(w, http.StatusInternalServerError)
}
