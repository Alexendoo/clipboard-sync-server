package routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Alexendoo/sync/model"
)

type newUserRequest struct {
	Link []byte
}

type newUserResponse struct {
	User   *model.User
	Device *model.Device
}

// Register creates a new User and provision an initial Device
func Register(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		badRequest(w)
		return
	}

	var req newUserRequest
	err = json.Unmarshal(bytes, &req)
	if err != nil || req.Link == nil {
		badRequest(w)
		return
	}

	user := model.NewUser()
	err = user.Save(db)
	if err != nil {
		serverError(w)
		return
	}

}
