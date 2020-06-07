package acme // import "taylz.io/http/route/acme"

import (
	"net/http"

	"taylz.io/http/route"
	"taylz.io/http/router"
)

// Thumbprint creates a new *mux.Route for the given file system path to use for stateless ACME challenges on path "/.well-known/acme-challenge/"
func Thumbprint(thumbprint string) *route.T {
	lencut := 28 // len("/.well-known/acme-challenge/")
	addthumb := "." + thumbprint
	return &route.T{
		Router: router.PathStarts("/.well-known/acme-challenge"),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if len(r.URL.Path) < lencut {
				w.Write([]byte("error: path too short"))
				return
			}
			match := r.URL.Path[lencut:]
			w.Write([]byte(match + addthumb))
		}),
	}
}
