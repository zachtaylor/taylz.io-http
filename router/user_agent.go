package router

import "net/http"

// UserAgent is a Router for matching User-Agent
type UserAgent string

// RouteHTTP matches the first chars of `Request.Header["User-Agent"`
func (ua UserAgent) RouteHTTP(r *http.Request) bool {
	header := r.Header.Get("User-Agent")
	lenua := len(ua)
	return len(header) >= lenua && header[:lenua] == string(ua)
}
func (ua UserAgent) isRouter() I { return ua }
