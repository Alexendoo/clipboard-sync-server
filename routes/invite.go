package routes

import (
	"io"
	"net/http"

	"sync"

	"log"

	"github.com/gorilla/mux"
)

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
			make(chan *http.Request),
			make(chan http.ResponseWriter),
			make(chan *http.Request),
			make(chan http.ResponseWriter),
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

type invite struct {
	wg sync.WaitGroup

	srcReq chan *http.Request
	srcRes chan http.ResponseWriter

	destReq chan *http.Request
	destRes chan http.ResponseWriter
}

var invites = NewInvites()

func InviteGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	key := vars["key"]
	deviceType := vars["device"]

	inv := invites.Get(key)
	defer inv.wg.Done()
	defer log.Println("req done")

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

	inv := invites.Get(key)
	defer inv.wg.Done()
	defer log.Println("post done")

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
