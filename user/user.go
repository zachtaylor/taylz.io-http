package user

import (
	"taylz.io/http/session"
	"taylz.io/types"
)

type T struct {
	settings *Settings
	session  *session.T
	socks    *types.SetString
}

// Name returns Session name, if available
func (t *T) Name() string {
	if t.session != nil {
		return t.session.Name()
	}
	return ""
}

// WriteJSON is a macro for Write() with json encoding
func (t *T) WriteJSON(dict types.Dict) {
	t.Write(types.NewBytesString(types.NewStringDict(dict)))
}

func (t *T) Write(data types.Bytes) {
	for _, socketid := range t.socks.SliceString() {
		if socket := t.settings.Sockets.Get(socketid); socket != nil {
			socket.Write(data)
		}
	}
}

// AddSocketID adds a socket id to the user message dispatcher
func (t *T) AddSocketID(socketid string) { t.socks.Add(socketid) }

// RemoveSocketID removes a socket id from the user message dispatcher
func (t *T) RemoveSocketID(socketid string) { t.socks.Remove(socketid) }

func (t *T) destroy() {
	t.settings = nil
	t.session = nil
	t.socks = nil
}
