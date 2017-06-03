package routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Alexendoo/sync/model"
)

type newUserRequest struct {
	Link      []byte
	FCMToken  string
	Signature []byte
	PubKey    []byte
}

type newUserResponse struct {
	User   *model.User
	Device *model.Device
}

// RegisterUser creates a new User and provision an initial Device
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		badRequest(w)
		return
	}

	var req newUserRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		badRequest(w)
		return
	}

	user := model.NewUser()
	err = user.Save(db)
	if err != nil {
		serverError(w)
		return
	}

	device := model.NewDevice(req.PubKey, "foo", req.FCMToken, user.ID)
	err = device.Save(db)
	if err != nil {
		serverError(w)
		user.Delete(db)
		return
	}

	linkAdded := AddLink(w, user.ID, req.Link, req.Signature, req.PubKey)
	if !linkAdded {
		device.Delete(db)
		user.Delete(db)
		return
	}

	resp, _ := json.Marshal(&newUserResponse{user, device})
	w.Write(resp)
}
