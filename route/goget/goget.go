package goget // import "taylz.io/http/route/goget"

import (
	"net/http"
	"text/template"

	"taylz.io/http/route"
	"taylz.io/http/router"
)

// Domain creates a new *mux.Route which handles go get style challenges for the given domain name
func Domain(domain string) *route.T {
	return &route.T{
		Router:  router.UserAgent("Go-http-client"),
		Handler: NewHandler(domain),
	}
}

const tpls = `<html>
	<meta name="go-import" content="{{.Host}}/{{.Package}} git https://{{.Host}}/{{.Package}}">
	<meta name="go-source" content="{{.Host}}/{{.Package}} https://{{.Host}}/{{.Package}} https://{{.Host}}/{{.Package}}/tree/master{/dir} https://{{.Host}}/{{.Package}}/tree/master{/dir}/{file}#L{line}">
</html>`

type tpld struct {
	Host    string
	Package string
}

// NewHandler returns a http.Handler that writes data for the go tool to find code using go get
//
// Note that go requires "git+https://{{host}}/" to work
func NewHandler(host string) http.Handler {
	t := template.Must(template.New("").Parse(tpls))
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pkg := r.RequestURI[1 : len(r.RequestURI)-len("?go-get=1")]
		t.Execute(w, tpld{
			Host:    host,
			Package: pkg,
		})
	})
}
