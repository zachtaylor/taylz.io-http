package server

import (
	"net/http"

	"taylz.io/http/session"
	"taylz.io/http/user"
	"taylz.io/http/websocket"
	"taylz.io/keygen"
	"taylz.io/types"
	"taylz.io/z/charset"
)

type T struct {
	mux      *Mux
	wsmux    *websocket.Mux
	Sessions session.Storer
	Sockets  websocket.Storer
	Users    user.Storer
}

func New() *T {
	sessions := session.NewServer(session.SettingsDefault, session.NewCache())
	keepalive := 30 * types.Second
	wsmux := &websocket.Mux{}
	sockets := websocket.NewServer(
		websocket.Settings{
			KeepAlive: &keepalive,
			Keygen: &keygen.Settings{
				KeySize: 4,
				CharSet: charset.AlphaCapitalNumeric,
				Rand:    types.NewRand(types.NewSeeder(types.NewTime().UnixNano())),
			},
		},
		websocket.NewCache(),
		wsmux,
	)
	users := user.NewServer(
		user.Settings{
			Sessions: sessions,
			Sockets:  sockets,
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
