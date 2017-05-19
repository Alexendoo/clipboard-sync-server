package routes

import (
	"database/sql"
	"net/http"

	"github.com/Alexendoo/sync/model"
	"github.com/gorilla/mux"
)

// Handler for all the application routes
func Handler(db *sql.DB) http.Handler {
	router := mux.NewRouter()

	devices := model.NewDeviceStore(db)
	users := model.NewUserStore(db)

	router.HandleFunc("/register", Register(users, devices)).
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
