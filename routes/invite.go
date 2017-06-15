package routes

import (
	"context"
	"io"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

// NewInviteHandler returns a new instance of InviteHandler
func NewInviteHandler() *InviteHandler {
	upgrader := &websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	connections := make(map[string]*pending)

	return &InviteHandler{
		upgrader:    upgrader,
		connections: connections,
	}
}

type pending struct {
	w      http.ResponseWriter
	r      *http.Request
	cancel context.CancelFunc
}

func newPending(w http.ResponseWriter, r *http.Request) *pending {
	parentCtx := r.Context()
	ctx, cancel := context.WithCancel(parentCtx)

	return &pending{
		w:      w,
		r:      r.WithContext(ctx),
		cancel: cancel,
	}
}

// InviteHandler rendezvous client pairs at an invite URL
type InviteHandler struct {
	sync.RWMutex
	upgrader    *websocket.Upgrader
	connections map[string]*pending
}

func (h *InviteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]

	if !websocket.IsWebSocketUpgrade(r) {
		notFound(w)
		return
	}

	h.RLock()
	pending, ok := h.connections[key]
	h.RUnlock()

	if ok {
		h.connect(w, r, pending)
		return
	}

	pending = newPending(w, r)

	h.Lock()
	h.connections[key] = pending
	h.Unlock()

	<-pending.r.Context().Done()

	h.Lock()
	delete(h.connections, key)
	h.Unlock()
}

// connect upgrades both (w, r) and (pending.w, pending.r) to websockets and
// connects them bidrectionally
func (h *InviteHandler) connect(w http.ResponseWriter, r *http.Request, pending *pending) {
	src, err := h.upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	dest, err := h.upgrader.Upgrade(pending.w, pending.r, nil)
	if err != nil {
		src.Close()
		return
	}

	select {
	case <-pipeWebsocket(src, dest):
	case <-pipeWebsocket(dest, src):
	}

	pending.cancel()
	src.Close()
	dest.Close()
}

func pipeWebsocket(src, dest *websocket.Conn) <-chan interface{} {
	done := make(chan interface{})
	go func() {
		for {
			messageType, r, err := src.NextReader()
			if err != nil {
				break
			}

			w, err := dest.NextWriter(messageType)
			if err != nil {
				break
			}

			if _, err := io.Copy(w, r); err != nil {
				break
			}

			if err := w.Close(); err != nil {
				break
			}
		}

		close(done)
	}()
	return done
}
