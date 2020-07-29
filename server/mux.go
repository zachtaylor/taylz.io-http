package server

import (
	"net/http"

	"taylz.io/http/route"
	"taylz.io/http/router"
)

// Mux is slice of *T that implements http.Handler
type Mux []*route.T

// Add adds Route to this Mux
func (mux *Mux) Add(t *route.T) {
	*mux = append(*mux, t)
}

// Route is a macro for Add(T{})
func (mux *Mux) Route(r router.I, h http.Handler) {
	mux.Add(&route.T{
		Router:  r,
		Handler: h,
	})
}

// Handle is a macro for Route(router.Path(route), h)
func (mux *Mux) Handle(route string, h http.Handler) {
	mux.Route(router.Path(route), h)
}

// ServeHTTP dispatched the request to the internal route
func (mux *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var handler http.Handler
	for _, route := range *mux {
		if route.Router.RouteHTTP(r) {
			handler = route.Handler
			break
		}
	}
	if handler != nil {
		handler.ServeHTTP(w, r)
	}
}
