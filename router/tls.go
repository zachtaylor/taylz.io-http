package router

import "net/http"

// TLSOn satisfies `Router` by matching `Request.TLS` is non-nil
var TLSOn = Func(func(r *http.Request) bool {
	return r.TLS != nil
})

// TLSOff satisfies `Router` by matching `Request.TLS` is nil
var TLSOff = Func(func(r *http.Request) bool {
	return r.TLS == nil
})
