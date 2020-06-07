package handler

import "net/http"

// RedirectHTTPS is a Handler that always uses http.Redirect to direct a request to https
var RedirectHTTPS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://"+r.Host+r.URL.String(), http.StatusMovedPermanently)
})
