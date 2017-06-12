package routes

import (
	"net/http"

	"github.com/Alexendoo/clipboard-sync-server/messages"
	"github.com/golang/protobuf/proto"
)

// About exposes server information to clients
func About(w http.ResponseWriter, r *http.Request) {
	info := &messages.ServerInfo{
		SenderId: "303334042045",
		Version:  "0.1.0",
	}

	bytes, err := proto.Marshal(info)
	if err != nil {
		panic(err)
	}

	w.Write(bytes)
}
