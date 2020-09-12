package server

import (
	"net/http"

	"taylz.io/http/session"
	"taylz.io/http/user"
	"taylz.io/http/websocket"
)

type T struct {
	mux      *Mux
	wsmux    *websocket.Mux
	Sessions *session.Server
	Sockets  *websocket.Server
	Users    *user.Server
}

func New(sesset session.Settings, sckset websocket.Settings) *T {
	sessions := session.NewServer(sesset, session.NewCache())
	wsmux := &websocket.Mux{}
	sockets := websocket.NewServer(sckset, websocket.NewCache(), wsmux)
	users := user.NewServer(
		user.Settings{
			Sessions: sessions,
			Sockets:  sockets.Cache,
		},
		user.NewCache(),
	)
	return &T{
		mux:      &Mux{},
		wsmux:    wsmux,
		Sessions: sessions,
		Sockets:  sockets,
		Users:    users,
	}
}

func (t *T) Mux() *Mux             { return t.mux }
func (t *T) WSMux() *websocket.Mux { return t.wsmux }

func (t *T) ServeHTTP(w http.ResponseWriter, r *http.Request) { t.mux.ServeHTTP(w, r) }
