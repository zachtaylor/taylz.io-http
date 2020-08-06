package user

import (
	"taylz.io/http/session"
	"taylz.io/http/websocket"
)

// Server is a user service
type Server struct {
	Settings Settings
	Storer
	userSocket  map[string]string
	userSession map[string]string
}

// NewServer creates a user server
func NewServer(settings Settings, store Storer) (server *Server) {
	server = &Server{
		Settings:    settings,
		Storer:      store,
		userSocket:  make(map[string]string),
		userSession: make(map[string]string),
	}

	settings.Sessions.Observe(server.onSession)
	settings.Sockets.Observe(server.onWebsocket)

	return
}

func (s *Server) onSession(id string, session *session.T) {
	s.Sync(func(get Getter, set Setter) {
		if session == nil {
			set(s.userSession[id], nil)
			delete(s.userSession, id)
		} else {
			s.userSession[id] = session.Name()
			set(session.Name(), New(&s.Settings, session))
		}
	})
}

func (s *Server) onWebsocket(id string, ws *websocket.T) {
	if ws == nil {
		if user := s.Get(s.userSocket[id]); user != nil {
			user.RemoveSocketID(id)
		}
		delete(s.userSocket, id)
		return
	}

	session := s.Settings.Sessions.RequestSessionCookie(ws.Conn.Request())
	if session == nil {
		return
	}
	s.AddUser(session, ws)
}

// GetUser returns the user associated with the websocket id
func (s *Server) GetUser(ws *websocket.T) *T { return s.Get(s.userSocket[ws.ID()]) }

// AddUser links a websocket to a user manually
func (s *Server) AddUser(session *session.T, ws *websocket.T) {
	if u := s.Get(session.Name()); u != nil {
		s.userSocket[ws.ID()] = session.Name()
		u.AddSocketID(ws.ID())
	}
}
