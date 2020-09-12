package user

import (
	"taylz.io/http/session"
	"taylz.io/http/websocket"
)

// Server is a user service
type Server struct {
	Settings Settings
	*Cache
}

// NewServer creates a user server
func NewServer(settings Settings, cache *Cache) (server *Server) {
	server = &Server{
		Settings: settings,
		Cache:    cache,
	}

	settings.Sessions.Observe(server.onSession)
	settings.Sockets.Observe(server.onWebsocket)

	return
}

func (s *Server) onSession(id string, oldSession, newSession *session.T) {
	s.Sync(func(get CacheGetter, set CacheSetter) {
		if newSession == nil {
			set(oldSession.Name(), nil)
		} else {
			set(newSession.Name(), New(&s.Settings, newSession))
		}
	})
}

func (s *Server) onWebsocket(id string, oldWS, newWS *websocket.T) {
	if newWS != nil {
		if session := s.Settings.Sessions.RequestSessionCookie(newWS.Request()); session != nil {
			s.Pair(session, newWS)
		}
	} else if user := s.GetSocket(oldWS); user != nil {
		user.RemoveSocketID(id)
	}
}

// GetSocket returns the user associated with the websocket id
func (s *Server) GetSocket(ws *websocket.T) *T {
	if session := s.Settings.Sessions.RequestSessionCookie(ws.Request()); session != nil {
		return s.Get(session.Name())
	}

	return nil
}

// Pair links a websocket to a user manually
func (s *Server) Pair(session *session.T, ws *websocket.T) {
	if user := s.Get(session.Name()); user != nil {
		user.AddSocketID(ws.ID())
	}
}
