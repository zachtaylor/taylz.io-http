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
func (t *T) Sockets() types.SliceString { return t.socks.Slice() }

// Message calls Write using websocket.Transport data format
func (t *T) Message(uri string, data types.Dict) { t.Write(websocket.Transport(uri, data)) }

// Write writes the buffer to all sockets
func (t *T) Write(bytes types.Bytes) {
	for _, k := range t.Sockets() {
		if ws := t.settings.Sockets.Get(k); ws != nil {
			ws.Write(bytes)
		}
	}
}
