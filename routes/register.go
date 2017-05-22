package routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Alexendoo/sync/model"
)

type registrationRequest struct {
	Name  string
	Token string
}

type registrationResponse struct {
	User   *model.User
	Device *model.Device
}

// Register a new User and provision an initial Device
func Register(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		badRequest(w)
		return
	}

	var req registrationRequest
	err = json.Unmarshal(bytes, &req)
	if err != nil || req.Name == "" || req.Token == "" {
		badRequest(w)
		return
	}

	user := model.NewUser()
	device := model.NewDevice(req.Name, req.Token, user.ID)

	err = user.Save(db)
	if err != nil {
		serverError(w)
		return
	}
	err = device.Save(db)
	if err != nil {
		user.Delete(db)
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
