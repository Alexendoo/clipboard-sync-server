package routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Alexendoo/sync/model"
)

type newUserRequest struct {
	Name  string
	Token string
}

type newUserResponse struct {
	User   *model.User
	Device *model.Device
}

// RegisterUser creates a new User and provision an initial Device
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		badRequest(w)
		return
	}

	var req newUserRequest
	err = json.Unmarshal(bytes, &req)
	if err != nil || empty(req.Name, req.Token) {
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

	json, err := json.Marshal(&newUserResponse{user, device})
	if err != nil {
		serverError(w)
		return
	}

	w.Write(json)
}

type newDeviceRequest struct {
	Name   string
	Token  string
	UserID string
}

type newDeviceResponse struct {
	Device *model.Device
}

// RegisterDevice attaches a new Device to an existing User
func RegisterDevice(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		badRequest(w)
		return
	}

	var req newDeviceRequest
	err = json.Unmarshal(bytes, &req)
	if err != nil || empty(req.Name, req.Token, req.UserID) {
		badRequest(w)
		return
	}

	ok, err := model.UserExists(db, req.UserID)
	if !ok || err != nil {
		badRequest(w)
		return
	}

	device := model.NewDevice(req.Name, req.Token, req.UserID)
	err = device.Save(db)
	if err != nil {
		serverError(w)
		return
	}

	json, err := json.Marshal(&newDeviceResponse{
		Device: device,
	})
	if err != nil {
		serverError(w)
		return
	}

	w.Write(json)
}
