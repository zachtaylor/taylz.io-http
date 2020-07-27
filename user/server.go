package user

import (
	"taylz.io/http/session"
	"taylz.io/http/websocket"
)

// Server is a user service
type Server struct {
	Settings Settings
	Storer
	live map[string]string
}

// NewServer creates a user server
func NewServer(settings Settings, store Storer) (server *Server) {
	server = &Server{
		Settings: settings,
		Storer:   store,
		live:     make(map[string]string),
	}

	settings.Sessions.Observe(server.onSession)
	settings.Sockets.Observe(server.onWebsocket)

	return
}

func (s *Server) onSession(id string, session *session.T) {
	s.Sync(func(dat map[string]*T) {
		if session == nil {
			delete(dat, id)
		} else if dat[id] == nil {
			dat[id] = New(session)
		}
	})
}

func (s *Server) onWebsocket(id string, ws *websocket.T) {
	if ws == nil {
		if user := s.Get(s.live[id]); user != nil {
			user.RemoveSocketID(id)
		}
		delete(s.live, id)
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
	if u := s.Get(session.Name()); u != nil {
		s.live[ws.ID()] = session.Name()
		u.AddSocketID(ws.ID())
	}
}
