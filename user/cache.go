package user

import (
	"taylz.io/http/session"
	"taylz.io/types"
)

// Cache is a user store
type Cache struct {
	settings *Settings
	cache    map[string]*T
	lock     types.Mutex
}

// Get returns the user associated to the name, if available
func (c *Cache) Get(name string) *T {
	return c.cache[name]
}

// User returns a user for the session, creates it if necessary
func (c *Cache) User(session *session.T) (t *T) {
	c.lock.Lock()
	if t = c.cache[session.Name()]; t == nil {
		t = &T{
			settings: c.settings,
			session:  session,
			socks:    types.NewSetString(),
		}
		go c.wait(t)
		c.cache[session.Name()] = t
	}
	c.lock.Unlock()
	return
}

func (c *Cache) wait(t *T) {
	<-t.session.Done()
	c.lock.Lock()
	delete(c.cache, t.session.Name())
	c.lock.Unlock()
	t.destroy()
}
