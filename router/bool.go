package router

import "net/http"

// Bool satisfies `Router` by returning a constant
type Bool bool

// RouteHTTP satisfies `Router` by returning a constant
func (b Bool) RouteHTTP(_ *http.Request) bool { return bool(b) }
func (b Bool) isRouter() I                    { return b }

// BoolTrue is a `Router` that always returns true
var BoolTrue Bool = true
