package routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Alexendoo/sync/model"
)

type registrationRequest struct {
	Name string
}

type registrationResponse struct {
	User   *model.User
	Device *model.Device
}

// Register a new User and provision an initial Device
func Register(users *model.UserStore, devices *model.DeviceStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			badRequest(w)
			return
		}

		var req registrationRequest
		err = json.Unmarshal(bytes, &req)
		if err != nil || req.Name == "" {
			badRequest(w)
			return
		}

		user := model.NewUser()
		device := model.NewDevice(req.Name, user)

		err = users.Insert(user)
		if err != nil {
			serverError(w)
			return
		}
		err = devices.Insert(device)
		if err != nil {
			users.Delete(user)
			serverError(w)
			return
		}

		json, err := json.Marshal(&registrationResponse{user, device})
		if err != nil {
			serverError(w)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	}
}
