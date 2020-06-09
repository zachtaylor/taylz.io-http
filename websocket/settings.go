package websocket

import (
	"taylz.io/http/session"
	"taylz.io/keygen"
	"taylz.io/types"
)

// Settings is configuration for Cache behavior management
type Settings struct {
	// UsePing enables keep-alive encouragement at given speed
	UsePing *types.Duration
	// KeyGener controls new Websocket IDs
	Keygen keygen.I
	// Sessions controls association of Websocket with Session
	Sessions *session.Cache
	// Server controls the API available for received messages
	Handler Handler
}
