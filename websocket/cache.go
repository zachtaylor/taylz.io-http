package websocket

import (
	"net/http"

	"golang.org/x/net/websocket"
	"taylz.io/types"
)

// Cache manages websocket connections
type Cache struct {
	settings Settings
	cache    map[string]*T
	lock     types.Mutex // guard cache
}

// NewCache builds a Cache, required for internals
func NewCache(settings Settings) *Cache {
	return &Cache{
		settings: settings,
		cache:    make(map[string]*T),
	}
}

// Count returns the number of open sockets
func (c *Cache) Count() int {
	return len(c.cache)
}

// Has implements KeyStorer
func (c *Cache) Has(id string) (ok bool) {
	_, ok = c.cache[id]
	return
}

// Get returns the websocket for the given key
func (c *Cache) Get(id string) *T {
	return c.cache[id]
}

// Remove deletes a socket ID from the cache
func (c *Cache) Remove(id string) {
	c.lock.Lock()
	delete(c.cache, id)
	c.lock.Unlock()
}

// Keys creates a new slice of the active connected keys
func (c *Cache) Keys() []string {
	c.lock.Lock()
	keys := make([]string, len(c.cache))
	i := 0
	for k := range c.cache {
		keys[i] = k
		i++
	}
	c.lock.Unlock()
	return keys
}

// Handler returns a new http.Handler that adds upgrades requests to add Websockets to this cache
func (c *Cache) Handler() http.Handler {
	return websocket.Handler(c.connect)
}

func (c *Cache) connect(conn *websocket.Conn) {
	c.lock.Lock()
	var id string
	for ok := true; ok; _, ok = c.cache[id] {
		id = c.settings.Keygen.Keygen()
	}
	ws := New(id, conn)
	c.cache[id] = ws
	c.lock.Unlock()

	if session := c.settings.Sessions.RequestSessionCookie(conn.Request()); session == nil {
	} else { // has session
		ws.Session = session
	}

	if c.settings.UsePing == nil {
		c.watch(ws)
	} else {
		c.watchWithMonitor(ws)
	}

	if ws.Session != nil {
		ws.Session = nil
	}

	c.Remove(ws.id)
}

// watchWithMonitor performs socket i/o and sends json when lonely
func (c *Cache) watchWithMonitor(ws *T) {
	for lonelyTimer, resetCD := types.NewTimer(*c.settings.UsePing), types.NewTime(); ; {
		select {
		case <-lonelyTimer.C:
			ws.Write(wsLonely)
			lonelyTimer.Reset(*c.settings.UsePing)
		case buff := <-ws.send: // write to client
			// lonelyTimer.Reset on write effect has 1 sec cooldown
			if now := types.NewTime(); now.Sub(resetCD) > types.Second {
				if !lonelyTimer.Stop() {
					<-lonelyTimer.C
				}
				lonelyTimer.Reset(*c.settings.UsePing)
				resetCD = now
			}
			if err := websocket.Message.Send(ws.conn, buff); err != nil {
				if !lonelyTimer.Stop() {
					<-lonelyTimer.C
				}
				go wsDrainMessageChan(ws.recv)
				ws.Close()
				return
			}
		case msg := <-ws.recv: // read from client
			if msg == nil {
				if !lonelyTimer.Stop() {
					<-lonelyTimer.C
				}
				ws.Close()
				return
			}
			// lonelyTimer.Reset on read effect has 1 sec cooldown
			if now := types.NewTime(); now.Sub(resetCD) > types.Second {
				if !lonelyTimer.Stop() {
					<-lonelyTimer.C
				}
				lonelyTimer.Reset(*c.settings.UsePing)
				resetCD = now
			}
			go c.settings.Handler.ServeWS(ws, msg)
		}
	}
}

// watch performs socket i/o without additional helpers
func (c *Cache) watch(ws *T) {
	for {
		select {
		case buff := <-ws.send: // write to client
			if err := websocket.Message.Send(ws.conn, buff); err != nil {
				go wsDrainMessageChan(ws.recv)
				ws.Close()
				return
			}
		case msg := <-ws.recv: // read from client
			if msg == nil {
				ws.Close()
				return
			}
			go c.settings.Handler.ServeWS(ws, msg)
		}
	}
}
