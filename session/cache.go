package session

import "sync"

// Cache manages Sessions
type Cache struct {
	settings Settings
	cache    map[string]*T
	lock     sync.Mutex // guards cache write
}

// NewCache creates a new Sessions cache
func NewCache(settings Settings) *Cache {
	c := &Cache{
		settings: settings,
		cache:    make(map[string]*T),
	}
	return c
}

// Count returns number of active Sessions
func (c *Cache) Count() int {
	return len(c.cache)
}

// Has implements KeyStorer
func (c *Cache) Has(id string) (ok bool) {
	_, ok = c.cache[id]
	return
}

// Get returns a Session by id, if any
func (c *Cache) Get(id string) *T {
	return c.cache[id]
}

// Find returns a Session by name, if any
func (c *Cache) Find(name string) (session *T) {
	c.lock.Lock() // guards cache write
	for _, s := range c.cache {
		if name == s.name {
			session = s
			break
		}
	}
	c.lock.Unlock()
	return
}

// Grant returns a new Session granted to the username
//
// This is the canonical way to create a Session
func (c *Cache) Grant(name string) *T {
	c.lock.Lock() // guards cache write
	var id string
	for ok := false; !ok; _, ok = c.cache[id] {
		id = c.settings.Keygen.Keygen()
	}
	session := &T{
		id:    id,
		name:  name,
		in:    make(chan bool),
		done:  make(chan bool),
		socks: make([]string, 0),
	}
	c.cache[id] = session
	c.lock.Unlock()
	go c.watch(session)
	return session
}

// watch runs session.watch, then remove from the cache
func (c *Cache) watch(session *T) {
	session.watch(c.settings.Lifetime)
	c.lock.Lock() // guards cache write
	delete(c.cache, session.id)
	c.lock.Unlock()
}
