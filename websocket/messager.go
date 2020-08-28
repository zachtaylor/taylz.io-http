package websocket

import "taylz.io/types"

// Messager is a target for sending data to
type Messager interface {
	// Message writes to the client
	Message(string, types.Dict)
}
