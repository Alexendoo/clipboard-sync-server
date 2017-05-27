package routes

import (
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

// Invites is a list of invites that exist for the lifetime of any requests to the
// invite key
type Invites struct {
	sync.RWMutex
	invites map[string]*invite
}

func NewInvites() *Invites {
	return &Invites{
		invites: make(map[string]*invite),
	}
}

// Get or create an invite
func (i *Invites) Get(key string) *invite {
	i.RLock()
	inv, ok := i.invites[key]
	i.RUnlock()

	if !ok {

		inv = &invite{
			sync.WaitGroup{},
			newInviteConn(),
			newInviteConn(),
		}

		i.Lock()
		i.invites[key] = inv
		i.Unlock()

		inv.wg.Add(1)

		go func() {
			inv.wg.Wait()
			log.Printf("invite key freed: %v", key)
			delete(i.invites, key)
		}()
	} else {
		inv.wg.Add(1)
	}

	return inv
}

type inviteConn struct {
	res  chan http.ResponseWriter
	done chan interface{}
}

func newInviteConn() *inviteConn {
	return &inviteConn{
		make(chan http.ResponseWriter),
		make(chan interface{}),
	}
}

type invite struct {
	wg sync.WaitGroup

	src  *inviteConn
	dest *inviteConn
}

var invites = NewInvites()

func InviteGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	key := vars["key"]
	deviceType := vars["device"]

	inv := invites.Get(key)
	defer inv.wg.Done()

	var device *inviteConn
	if deviceType == "src" {
		device = inv.src
	} else {
		device = inv.dest
	}

	device.res <- w
	<-device.done
}

func InvitePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	key := vars["key"]
	deviceType := vars["device"]

	inv := invites.Get(key)
	defer inv.wg.Done()

	var device *inviteConn
	if deviceType == "src" {
		device = inv.dest
	} else {
		device = inv.src
	}

	res := <-device.res

	res.Header().Set("Content-Length", r.Header.Get("Content-Length"))

	_, err := io.Copy(res, r.Body)
	device.done <- r

	if err != nil {
		serverError(w)
		return
	}
}
