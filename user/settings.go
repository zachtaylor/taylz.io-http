package user

import "taylz.io/http/websocket"

// Settings is configuration for user connections
type Settings struct {
	// Sockets cache pointer
	Sockets *websocket.Cache
}
