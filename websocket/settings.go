package websocket

import (
	"taylz.io/keygen"
	"taylz.io/types"
)

// Settings is configuration for websockets
type Settings struct {
	// KeepAlive enables keep-alive encouragement at given speed
	KeepAlive *types.Duration
	// Keygen controls new Websocket IDs
	Keygen keygen.I
}
