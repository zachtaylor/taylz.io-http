package session

import "net/http"

// Server is a session manager
type Server struct {
	Settings Settings
	*Cache
}

// NewServer creates a session server
func NewServer(settings Settings, cache *Cache) *Server {
	return &Server{
		Settings: settings,
		Cache:    cache,
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

// WriteSessionCookie writes the session cookie to the ResponseWriter
func (s *Server) WriteSessionCookie(w http.ResponseWriter, session *T) {
	s.writeSessionCookiePart(w, s.Settings.CookieID+"="+session.id+"; Path=/; ")
}

// WriteSessionCookieExpired writes an expired session cookie to the ResponseWriter
func (s *Server) WriteSessionCookieExpired(w http.ResponseWriter) {
	s.writeSessionCookiePart(w, s.Settings.CookieID+"=; Path=/; Expires==Thu, 01 Jan 1970 00:00:00 GMT; ")
}

func (s *Server) writeSessionCookiePart(w http.ResponseWriter, part string) {
	if s.Settings.Secure {
		part += "Secure; "
	}
	if s.Settings.Strict {
		part += "SameSite=Strict;"
	} else {
		part += "SameSite=Lax;"
	}
	w.Header().Set("Set-Cookie", part)

}

// Grant returns a new Session granted to the username
func (s *Server) Grant(name string) *T {
	return New(name, s.Cache, s.Settings.Keygen, s.Settings.Lifetime)
}
