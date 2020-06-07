package router

import "net/http"

// I is a routing interface
type I interface {
	RouteHTTP(*http.Request) bool
}
