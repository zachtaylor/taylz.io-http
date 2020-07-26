package user

import (
	"taylz.io/http/session"
	"taylz.io/types"
)

// T is a user
type T struct {
	Session *session.T
	socks   *types.SetString
}

// New returns a User with the Session
func New(session *session.T) *T {
	return &T{
		Session: session,
		socks:   types.NewSetString(),
	}
}

// AddSocketID adds a socket id to the user
func (t *T) AddSocketID(socketid string) { t.socks.Add(socketid) }

// RemoveSocketID removes a socket id from the user
func (t *T) RemoveSocketID(socketid string) { t.socks.Remove(socketid) }

// Sockets returns the socket ids linked with the user
func (t *T) Sockets() types.SliceString { return t.socks.SliceString() }
