package session

import (
	"time"

	"taylz.io/keygen"
)

// T is a Session
type T struct {
	id   string
	name string
	in   chan bool
	done chan bool
}

// New creates a Session
func New(name string, store Storer, keygen keygen.I, lifetime time.Duration) (session *T) {
	store.Sync(func(dat map[string]*T) {
		var id string
		for ok := true; ok; _, ok = dat[id] {
			id = keygen.New()
		}
		session = &T{
			id:   id,
			name: name,
			in:   make(chan bool),
			done: make(chan bool),
		}
		dat[id] = session
		go watch(store, session, lifetime)
	})
	return
}
func watch(store Storer, session *T, lifetime time.Duration) {
	session.watch(lifetime)
	store.Remove(session.ID())
}

// ID returns the Session ID
func (session *T) ID() string { return session.id }

// Name returns the name of this Session
func (session *T) Name() string { return session.name }

// Refresh sends a refresh signal
func (session *T) Refresh() { go session.send(true) }

// Close sends a close signal
func (session *T) Close() { go session.send(false) }

// Done returns the observe channel, or nil if the Session is already closed
func (session *T) Done() <-chan bool { return session.done }

// send controls the input to session lifetime, renew or expire
func (session *T) send(ok bool) { session.in <- ok }

func (session *T) String() string {
	if session == nil {
		return "nil"
	}
	return "Session#" + session.id
}

// watch monitors the session duration, and can be renewed for the same duration, or stopped
func (session *T) watch(d time.Duration) {
	defer session.close()
	timer := time.NewTimer(d)
	for {
		select {
		case ok := <-session.in:
			if !timer.Stop() {
				<-timer.C // if can't stop, drain the channel
			}
			if !ok { // signal close
				return
			} // signal refresh
			timer.Reset(d)
		case <-timer.C:
			return
		}
	}
}

// close kills the Session
func (session *T) close() {
	close(session.in)
	close(session.done)
}
