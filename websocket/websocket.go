package websocket

import (
	"golang.org/x/net/websocket"
	"taylz.io/keygen"
	"taylz.io/types"
)

// T is a Websocket
type T struct {
	Conn *Conn
	id   string
	send chan types.Bytes
	recv <-chan *Message
	done chan bool
}

// Conn = websocket.Conn
type Conn = websocket.Conn

// Upgrader = websocket.Handler
type Upgrader = websocket.Handler

// New creates an initialized orphan Websocket
func New(conn *Conn, cache *Cache, keygen keygen.I) (ws *T) {
	cache.Sync(func(dat map[string]*T) {
		var id string
		for ok := true; ok; _, ok = dat[id] {
			id = keygen.New()
		}
		ws = &T{
			Conn: conn,
			id:   id,
			send: make(chan types.Bytes),
			recv: newChanMessage(conn),
			done: make(chan bool),
		}
		dat[id] = ws
		// go watch(cache, ws)
		// upgrader is responsible for goroutine
	})
	return
}

// ID returns the websocket ID
func (ws *T) ID() string { return ws.id }

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
		// close(ws.recv) managed by wsNewChanMessageConn
		ws.recv = nil
	}
}

// Message implements Messager
func (ws *T) Message(m *Message) {
	ws.Write(types.BytesString(types.StringDict(m.JSON())))
}
func (ws *T) isMessager() Messager { return ws }

var wsLonely = types.Bytes(`{"uri":"/ping"}`)

// newChanMessage creates a goroutine monitor using nextMessage
func newChanMessage(conn *Conn) <-chan *Message {
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
func nextMessage(conn *Conn) (*Message, error) {
	s, msg := "", &Message{}
	if err := websocket.Message.Receive(conn, &s); err != nil {
		return nil, err
	} else if err := types.DecodeJSON(types.BufferString(s), msg); err != nil {
		return nil, err
	}
	return msg, nil
}

// drainChanMessage waits to receive all messages, and returns when it reaches the end
func drainChanMessage(msgs <-chan *Message) {
	for {
		_, ok := <-msgs
		if !ok {
			return
		}
	}
}
