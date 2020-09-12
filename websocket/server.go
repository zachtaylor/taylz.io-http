package websocket

import "net/http"

// Server is a *Cache and Handler
type Server struct {
	Settings Settings
	*Cache
	Handler
}

// NewServer creates a websocket server
func NewServer(settings Settings, cache *Cache, h Handler) *Server {
	return &Server{
		Settings: settings,
		Cache:    cache,
		Handler:  h,
	}
}

// Upgrader returns a new http.Handler that adds upgrades request to add a Websocket to this cache
func (s *Server) Upgrader() http.Handler {
	return Upgrader(s.connect)
}
func (s *Server) connect(conn *Conn) {
	ws := New(conn, s.Cache, s.Settings.Keygen)
	if s.Settings.KeepAlive == nil {
		Watch(ws, s)
	} else {
		WatchWithMonitor(ws, *s.Settings.KeepAlive, s)
	}
	s.Remove(ws.id)
}
