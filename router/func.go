package router

import "net/http"

// Func satisfies `Router` by being a func
type Func func(*http.Request) bool

// RouteHTTP satisfies `Router` by calling the func
func (f Func) RouteHTTP(r *http.Request) bool { return f(r) }
func (f Func) isRouter() I                    { return f }
