package user

import (
	"taylz.io/http/session"
	"taylz.io/http/websocket"
)

// Settings is configuration for user cache connections
type Settings struct {
	// Sessions server pointer
	Sessions *session.Server
	// Sockets store
	Sockets websocket.Storer
}
