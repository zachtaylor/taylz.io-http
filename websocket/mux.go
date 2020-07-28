package websocket

// Mux is a slice of *Route that implements Handler
type Mux []*Route

// Add adds a Route to this Mux
func (mux *Mux) Add(r *Route) {
	*mux = append(*mux, r)
}

// Route is a macro for Add(Route{})
func (mux *Mux) Route(r Router, h Handler) {
	mux.Add(&Route{r, h})
}

// Handle is a macro for Route(RouterURI(route), h)
func (mux *Mux) Handle(route string, h Handler) {
	mux.Route(RouterURI(route), h)
}

// ServeWS satisfies Handler by routing to a matching *Route
func (mux *Mux) ServeWS(ws *T, m *Message) {
	var handler Handler
	for _, route := range *mux {
		if route.Router.RouteWS(m) {
			handler = route.Handler
			break
		}
	}
	if handler != nil {
		handler.ServeWS(ws, m)
	}
}
