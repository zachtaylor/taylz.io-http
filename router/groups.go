package router

import "net/http"

// And creates a `Router` group that returns true when all `Router` in the group return true
type And []I

// RouteHTTP satisfies `Router` by verifying all `Router` in the set return true
func (and And) RouteHTTP(r *http.Request) bool {
	for _, router := range and {
		if !router.RouteHTTP(r) {
			return false
		}
	}
	return true
}
func (and And) isRouter() I { return and }

// Or creates a `Router` group that returns true when any `Router` in the group returns true
type Or []I

// RouteHTTP satisfies Router by verifying any `Router` in the set returns true
func (or Or) RouteHTTP(r *http.Request) bool {
	for _, router := range or {
		if router.RouteHTTP(r) {
			return true
		}
	}
	return false
}
func (or Or) isRouter() I { return or }
