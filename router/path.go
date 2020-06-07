package router

import "net/http"

// Path satisfies `Router` by matching `Request.URL.Path` exactly
type Path string

// RouteHTTP satisfies `Router` by matching the request path exactly
func (path Path) RouteHTTP(r *http.Request) bool { return string(path) == r.URL.Path }
func (path Path) isRouter() I                    { return path }

// PathStarts satisfies `Router` by matching path starting with given prefix
type PathStarts string

// RouteHTTP satisfies `Router` by matching the path prefix
func (prefix PathStarts) RouteHTTP(r *http.Request) bool {
	lp := len(prefix)
	if len(r.URL.Path) < lp {
		return false
	}
	return string(prefix) == r.URL.Path[:lp]
}
func (prefix PathStarts) isRouter() I { return prefix }
