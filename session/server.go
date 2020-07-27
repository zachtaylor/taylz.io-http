package session

import "net/http"

// Server is a session manager
type Server struct {
	Settings Settings
	Storer
}

// NewServer creates a session server
func NewServer(settings Settings, store Storer) *Server {
	return &Server{
		Settings: settings,
		Storer:   store,
	}
}

// RequestSessionCookie returns Session associated to the Request via Session cookie
func (s *Server) RequestSessionCookie(r *http.Request) (session *T) {
	cookie, err := r.Cookie(s.Settings.CookieID)
	if err != nil {
		return nil
	}
	return s.Get(cookie.Value)
}

// Grant returns a new Session granted to the username
func (s *Server) Grant(name string) *T {
	return New(name, s, s.Settings.Keygen, s.Settings.Lifetime)
}
