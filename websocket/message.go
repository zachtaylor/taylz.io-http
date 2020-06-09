package websocket

import "taylz.io/types"

// Message is a simple messaging type
type Message struct {
	URI  string
	Data types.Dict
}

// NewMessage creates a Message
func NewMessage(uri string, data types.Dict) *Message {
	return &Message{
		URI:  uri,
		Data: data,
	}
}

// JSON returns json data
func (m *Message) JSON() types.Dict {
	if m == nil {
		return nil
	}
	return types.Dict{
		"uri":  m.URI,
		"data": m.Data,
	}
}

// Messager is an interface hook for Message targets
type Messager interface {
	// Message writes the Message to the receiver
	Message(*Message)
}
