package user

import (
	"taylz.io/http/session"
	"taylz.io/http/websocket"
	"taylz.io/types"
)

// T is a user
type T struct {
	settings *Settings
	Session  *session.T
	socks    *types.SetString
}

// New returns a User with the Session
func New(settings *Settings, session *session.T) *T {
	return &T{
		settings: settings,
		Session:  session,
		socks:    types.NewSetString(),
	}
}

// AddSocketID adds a socket id to the user
func (t *T) AddSocketID(socketid string) { t.socks.Add(socketid) }

// RemoveSocketID removes a socket id from the user
func (t *T) RemoveSocketID(socketid string) { t.socks.Remove(socketid) }

// Sockets returns the socket ids linked with the user
func (t *T) Sockets() types.SliceString { return t.socks.SliceString() }

// Message sends a message to all websockets
func (t *T) Message(m *websocket.Message) {
	for _, k := range t.Sockets() {
		if ws := t.settings.Sockets.Get(k); ws != nil {
			ws.Message(m)
		}
	}
}

func (t *T) Write(bytes types.Bytes) {
	for _, k := range t.Sockets() {
		if ws := t.settings.Sockets.Get(k); ws != nil {
			ws.Write(bytes)
		}
	}
}
