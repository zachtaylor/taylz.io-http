package router

import "net/http"

// Host satisfies `Router` by matching `Request.Host`
type Host string

// RouteHTTP satisfies `Router` by matching `Request.Host`
func (host Host) RouteHTTP(r *http.Request) bool { return string(host) == r.Host }
func (host Host) isRouter() I                    { return host }
