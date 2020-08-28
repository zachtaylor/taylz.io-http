package websocket

import (
	"net/http"

	"golang.org/x/net/websocket"
	"taylz.io/keygen"
	"taylz.io/types"
)

// Conn = websocket.Conn
type Conn = websocket.Conn

// Upgrader = websocket.Handler
type Upgrader = websocket.Handler

// Send calls the websocket send API
func Send(conn *Conn, bytes types.Bytes) error { return websocket.Message.Send(conn, bytes) }

// Receive calls the websocket receive API
func Receive(conn *Conn) (buf string, err error) {
	err = websocket.Message.Receive(conn, &buf)
	return
}

// T is a Websocket
type T struct {
	conn *Conn
	id   string
	send chan types.Bytes
	recv <-chan *Message
	done chan bool
}

//go:generate go-gengen -p=websocket -k=string -v=*T

// New creates a Websocket, which must be watched
func New(conn *Conn, store Storer, keygen keygen.I) (ws *T) {
	store.Sync(func(get Getter, set Setter) {
		var id string
		for ok := true; ok; ok = get(id) != nil {
			id = keygen.New()
		}
		ws = &T{
			conn: conn,
			id:   id,
			send: make(chan types.Bytes),
			recv: newChanMessage(conn),
			done: make(chan bool),
		}
		set(id, ws)
	})
	return
}

// ID returns the websocket ID
func (ws *T) ID() string { return ws.id }

// Request returns websocket establishment request
func (ws *T) Request() *http.Request { return ws.conn.Request() }

// Message calls Write using Message.JSON data format
func (ws *T) Message(uri string, data types.Dict) { ws.Write(Transport(uri, data)) }

// Write starts a goroutine to write bytes to to the socket API
func (ws *T) Write(bytes types.Bytes) { go ws.write(bytes) }
func (ws *T) write(bytes types.Bytes) { ws.send <- bytes }

// Send calls package-level Send with websocket.conn
func (ws *T) Send(bytes types.Bytes) error { return Send(ws.conn, bytes) }

// Close closes the observable channel
func (ws *T) Close() {
	if ws.done != nil {
		close(ws.send)
		close(ws.done)
		// close(ws.recv) managed by newChanMessage
		ws.done = nil
	}
}
