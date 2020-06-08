package mux

import (
	"net/http"

	"taylz.io/http/route"
)

// Mux is slice of *T that implements http.Handler
type Mux []*route.T

// Add adds Route to this Mux
func (mux *Mux) Add(t *T) {
	*mux = append(*mux, t)
}

// Route is a macro for Add(T{})
func (mux *Mux) Route(r I, h http.Handler) {
	mux.Add(&T{r, h})
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
