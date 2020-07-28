package websocket

// Route is a Router and Handler
type Route struct {
	Router  Router
	Handler Handler
}

// Router is used to route Messages
type Router interface {
	RouteWS(*Message) bool
}

// Handler is an interface hook for websocket API
type Handler interface {
	ServeWS(*T, *Message)
}

// HandlerFunc allows to make a func into a Handler
type HandlerFunc func(*T, *Message)

// ServeWS implements Handler by calling the func
func (f HandlerFunc) ServeWS(t *T, m *Message) { f(t, m) }
func (f HandlerFunc) isHandler() Handler       { return f }

// RouterURI creates a literal match check against Message.URI
type RouterURI string

// RouteWS implements Router by literally matching the Message.URI
func (r RouterURI) RouteWS(m *Message) bool { return string(r) == m.URI }
