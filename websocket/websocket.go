package websocket

import (
	"golang.org/x/net/websocket"
	"taylz.io/http/session"
	"taylz.io/types"
)

// T is a Websocket
type T struct {
	id      string
	conn    *websocket.Conn
	send    chan types.Bytes
	recv    <-chan *Message
	done    chan bool
	Session *session.T
}

// New creates an initialied orphan Websocket
func New(id string, conn *websocket.Conn) *T {
	return &T{
		id:   id,
		conn: conn,
		send: make(chan types.Bytes),
		recv: newChanMessage(conn),
		done: make(chan bool),
	}
}

// Write starts a goroutine to write bytes to to the socket API
func (ws *T) Write(buff types.Bytes) {
	go ws.write(buff)
}
func (ws *T) write(buff types.Bytes) {
	if ws.send != nil {
		ws.send <- buff
	}
}

// Close closes the observable channel
func (ws *T) Close() {
	if ws.done != nil {
		close(ws.send)
		ws.send = nil
		close(ws.done)
		ws.done = nil
		ws.recv = nil // close managed by wsNewChanMessageConn
	}
}

// Message implements Messager
func (ws *T) Message(m *Message) {
	ws.Write(types.BytesString(types.StringDict(m.JSON())))
}

var wsLonely = types.Bytes(`{"uri":"/ping"}`)

// newChanMessage creates a goroutine monitor using nextMessage
func newChanMessage(conn *websocket.Conn) <-chan *Message {
	msgs := make(chan *Message)
	go func() {
		for {
			if msg, err := nextMessage(conn); err == nil {
				msgs <- msg
			} else if err == types.EOF {
				break
			}
		}
		close(msgs)
	}()
	return msgs
}

// nextMessage synchronously reads a Message from the socket API
func nextMessage(conn *websocket.Conn) (*Message, error) {
	s, msg := "", &Message{}
	if err := websocket.Message.Receive(conn, &s); err != nil {
		return nil, err
	} else if err := types.DecodeJSON(types.BufferString(s), msg); err != nil {
		return nil, err
	}
	return msg, nil
}

// wsDrainMessageChan waits to receive all messages, and returns when it reaches the end
func wsDrainMessageChan(msgs <-chan *Message) {
	for {
		_, ok := <-msgs
		if !ok {
			return
		}
	}
}
