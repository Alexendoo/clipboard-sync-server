package routes

import (
	"database/sql"
	"net/http"

	"github.com/Alexendoo/sync/model"
	"github.com/gorilla/mux"
)

var (
	devices *model.DeviceStore
	users   *model.UserStore
)

// Handler for all the application routes
func Handler(db *sql.DB) http.Handler {
	router := mux.NewRouter()

	devices = model.NewDeviceStore(db).Debug()
	users = model.NewUserStore(db).Debug()

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
