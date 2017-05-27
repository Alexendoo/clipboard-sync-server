package routes

import (
	"net/http"

	"log"

	"io"

	"github.com/gorilla/mux"
)

type invite struct {
	srcReader *io.PipeReader
	srcWriter *io.PipeWriter

	destReader *io.PipeReader
	destWriter *io.PipeWriter
}

func newInvite() *invite {
	srcReader, destWriter := io.Pipe()
	destReader, srcWriter := io.Pipe()
	return &invite{
		srcReader,
		srcWriter,
		destReader,
		destWriter,
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
		_, err = io.Copy(w, inv.srcReader)
	} else {
		_, err = io.Copy(w, inv.destReader)
	}
	if err != nil {
		serverError(w)
	}

	log.Printf("inv: %#+v\n", inv)
	log.Printf("ok: %#+v\n", ok)
	log.Printf("key: %#+v\n", key)
	log.Printf("deviceType: %#+v\n", deviceType)
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
		_, err = io.Copy(inv.srcWriter, r.Body)
		inv.srcWriter.Close()
	} else {
		_, err = io.Copy(inv.destWriter, r.Body)
		inv.destWriter.Close()
	}
	if err != nil {
		serverError(w)
		return
	}
}
