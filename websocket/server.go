package websocket

import "net/http"

// Server is a *Cache and *Mux
type Server struct {
	Settings Settings
	Cache    *Cache
	Mux      *Mux
}

// NewServer creates a websocket server
func NewServer(settings Settings, cache *Cache, mux *Mux) *Server {
	return &Server{
		Settings: settings,
		Cache:    cache,
		Mux:      mux,
	}
}

// Upgrader returns a new http.Handler that adds upgrades request to add a Websocket to this cache
func (s *Server) Upgrader() http.Handler {
	return Upgrader(s.connect)
}
func (s *Server) connect(conn *Conn) {
	ws := New(conn, s.Cache, s.Settings.Keygen)
	s.Watch(ws)
	s.Cache.Remove(ws.id)
}

// Watch occupies the active goroutine with websocket monitor
func (s *Server) Watch(ws *T) {
	if s.Settings.KeepAlive == nil {
		Watch(ws, s.Mux)
	} else {
		WatchWithMonitor(ws, *s.Settings.KeepAlive, s.Mux)
	}
}
