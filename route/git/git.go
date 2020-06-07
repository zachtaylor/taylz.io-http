package git // import "taylz.io/http/route/git"

import (
	githttp "github.com/AaronO/go-git-http"
	"github.com/AaronO/go-git-http/auth"
	"taylz.io/http/route"
	"taylz.io/http/router"
)

// Path creates a new git route
func Path(path string) *route.T {
	return &route.T{
		Router:  router.UserAgent("git"),
		Handler: authNoPush(githttp.New(path)),
	}
}

// authNoPush is httpmuxgit middleware for restricting all git.Push requests
var authNoPush = auth.Authenticator(func(info auth.AuthInfo) (bool bool, error error) {
	return !info.Push, nil
})
