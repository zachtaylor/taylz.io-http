package user

import (
	"taylz.io/http/session"
	"taylz.io/http/websocket"
)

// Server is a *Cache and *Mux
type Server struct {
	Settings Settings
	Cache    *Cache
	live     map[string]string
}

// NewServer creates a user server
func NewServer(settings Settings, cache *Cache) (server *Server) {
	server = &Server{
		Settings: settings,
		Cache:    cache,
		live:     make(map[string]string),
	}

	settings.Sessions.Cache.Observe(server.onSession)
	settings.Sockets.Observe(server.onWebsocket)

	return
}

func (s *Server) onSession(id string, session *session.T) {
	if session == nil {
		s.Cache.Remove(id)
	} else if s.Cache.Get(id) == nil {
		s.Cache.Set(id, New(session))
	}
}

func (s *Server) onWebsocket(id string, ws *websocket.T) {
	if ws == nil {
		if user := s.Cache.Get(s.live[id]); user != nil {
			user.RemoveSocketID(id)
		}
		delete(s.live, id)
		return
	}

	if user := s.Cache.Get(s.live[id]); user != nil {
		user.AddSocketID(id)
		return
	}

	session := s.Settings.Sessions.RequestSessionCookie(ws.Conn.Request())
	if session == nil {
		return
	}
	s.AddUser(session, ws)
}

// AddUser links a websocket to a user manually
func (s *Server) AddUser(session *session.T, ws *websocket.T) {
	s.Cache.Get(session.Name()).AddSocketID(ws.ID())
	s.live[ws.ID()] = session.Name()
}
