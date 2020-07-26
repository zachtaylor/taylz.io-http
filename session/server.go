package session

import "net/http"

// Server is a session manager
type Server struct {
	Settings Settings
	Cache    *Cache
}

// RequestSessionCookie returns Session associated to the Request via Session cookie
func (s *Server) RequestSessionCookie(r *http.Request) (session *T) {
	cookie, err := r.Cookie(s.Settings.CookieID)
	if err != nil {
		return nil
	}
	return s.Cache.Get(cookie.Value)
}

// Grant returns a new Session granted to the username
func (s *Server) Grant(name string) *T {
	return New(name, s.Cache, s.Settings.Keygen, s.Settings.Lifetime)
}
