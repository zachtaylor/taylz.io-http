package session

import (
	"net/http"
	"time"
)

// T is a Session
type T struct {
	id    string
	name  string
	in    chan bool
	done  chan bool
	socks []string
}

// ID returns the Session ID
func (session *T) ID() string {
	return session.id
}

// Name returns the name of this Session
func (session *T) Name() string {
	return session.name
}

// Cookie returns a Cookie that encodes this Session
func (session *T) Cookie() *http.Cookie {
	return &http.Cookie{
		Name:  "SessionID",
		Value: session.id,
	}
}

// Refresh sends a refresh signal
func (session *T) Refresh() {
	go session.send(true)
}

// Close sends a close signal
func (session *T) Close() {
	go session.send(false)
}

// Done returns the observe channel, or nil if the Session is already closed
func (session *T) Done() <-chan bool {
	return session.done
}

func (session *T) String() string {
	if session == nil {
		return "nil"
	}
	return "Session#" + session.id
}

// send controls the input to session lifetime, renew or expire
func (session *T) send(ok bool) {
	session.in <- ok
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
