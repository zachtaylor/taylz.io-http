package route

import (
	"net/http"

	"taylz.io/http/router"
)

// T is a struct boxing Router and Handler
type T struct {
	Router  router.I
	Handler http.Handler
}

// RouteHTTP uses Router pointer
func (t *T) RouteHTTP(r *http.Request) bool { return t.Router.RouteHTTP(r) }
func (t *T) isRouter() router.I             { return t }

// ServeHTTP uses Handler pointer
func (t *T) ServeHTTP(w http.ResponseWriter, r *http.Request) { t.Handler.ServeHTTP(w, r) }
func (t *T) isHandler() http.Handler                          { return t }
