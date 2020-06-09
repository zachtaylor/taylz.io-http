package user

// import (
// 	"taylz.io/http/websocket"
// 	"taylz.io/types"
// )

// // Message is some content
// type Message struct {
// 	From string
// 	Time types.Time
// 	Body string
// }

// // Message returns a websocket.Message for writing to a websocket
// func (m *Message) Message(uri string) *websocket.Message {
// 	return &websocket.Message{
// 		URI:  uri,
// 		Data: m.JSON(),
// 	}
// }

// // JSON returns a Dict representing this Message
// func (m *Message) JSON() types.Dict {
// 	return types.Dict{
// 		"from": m.From,
// 		"time": m.Time.Format("2006-01-02 15:04:05"),
// 		"body": m.Body,
// 	}
// }
