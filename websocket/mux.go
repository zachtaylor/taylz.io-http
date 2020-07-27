package websocket

// Mux is a slice of *Route that implements Handler
type Mux []*Route

// Append adds a Route to this Mux
func (mux *Mux) Append(r *Route) {
	*mux = append(*mux, r)
}

// Route is a macro for Append(Route{})
func (mux *Mux) Route(r Router, h Handler) {
	mux.Append(&Route{r, h})
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
