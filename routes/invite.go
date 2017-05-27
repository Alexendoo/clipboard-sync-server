package routes

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type invite struct {
	srcReq chan *http.Request
	srcRes chan http.ResponseWriter

	destReq chan *http.Request
	destRes chan http.ResponseWriter
}

func newInvite() *invite {
	return &invite{
		make(chan *http.Request),
		make(chan http.ResponseWriter),
		make(chan *http.Request),
		make(chan http.ResponseWriter),
	}
}

var invites = make(map[string]*invite)

func InviteGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	key := vars["key"]
	deviceType := vars["device"]

	inv, ok := invites[key]

	if !ok {
		inv = newInvite()
		invites[key] = inv
	}

	var err error
	if deviceType == "src" {
		inv.srcRes <- w
		<-inv.srcReq
	} else {
		inv.destRes <- w
		<-inv.destReq
	}
	if err != nil {
		serverError(w)
	}
}

func InvitePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	key := vars["key"]
	deviceType := vars["device"]

	inv, ok := invites[key]

	if !ok {
		inv = newInvite()
		invites[key] = inv
	}

	var err error
	if deviceType == "src" {
		res := <-inv.destRes
		_, err = io.Copy(res, r.Body)
		inv.destReq <- r
	} else {
		res := <-inv.srcRes
		_, err = io.Copy(res, r.Body)
		inv.srcReq <- r
	}
	if err != nil {
		serverError(w)
		return
	}
}
