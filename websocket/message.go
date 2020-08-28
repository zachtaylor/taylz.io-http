package websocket

import "taylz.io/types"

// Message is a simple messaging type
type Message struct {
	URI  string
	Data types.Dict
}

// NewMessage creates a Message
func NewMessage(uri string, data types.Dict) *Message {
	return &Message{
		URI:  uri,
		Data: data,
	}
}

func newChanMessage(conn *Conn) <-chan *Message {
	msgs := make(chan *Message)
	go func() {
		for {
			if msg, err := receiveMessage(conn); err == nil {
				msgs <- msg
			} else if err == types.EOF {
				break
			}
		}
		close(msgs)
	}()
	return msgs
}

func receiveMessage(conn *Conn) (*Message, error) {
	buf, err := Receive(conn)
	if err != nil {
		return nil, err
	}
	msg := &Message{}
	if err = types.DecodeJSON(types.NewBufferString(buf), msg); err != nil {
		return nil, err
	}
	return msg, nil
}

func drainChanMessage(msgs <-chan *Message) {
	for ok := true; ok; _, ok = <-msgs {
	}
}
