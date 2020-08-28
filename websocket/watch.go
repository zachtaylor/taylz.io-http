package websocket

import "taylz.io/types"

var wsLonely = types.Bytes(`{"uri":"/ping"}`)

// WatchWithMonitor performs socket i/o and sends json when lonely
func WatchWithMonitor(ws *T, timeout types.Duration, handler Handler) {
	for lonelyTimer, resetCD := types.NewTimer(timeout), types.NewTime(); ; {
		select {
		case <-lonelyTimer.C:
			ws.Write(wsLonely)
			lonelyTimer.Reset(timeout)
		case buff := <-ws.send: // write to client
			if now := types.NewTime(); now.Sub(resetCD) > types.Second { // 1s cooldown
				if !lonelyTimer.Stop() {
					<-lonelyTimer.C
				}
				lonelyTimer.Reset(timeout)
				resetCD = now
			}
			if err := ws.Send(buff); err != nil {
				if !lonelyTimer.Stop() {
					<-lonelyTimer.C
				}
				go drainChanMessage(ws.recv)
				ws.Close()
				return
			}
		case msg := <-ws.recv: // read from client
			if msg == nil {
				if !lonelyTimer.Stop() {
					<-lonelyTimer.C
				}
				ws.Close()
				return
			}
			if now := types.NewTime(); now.Sub(resetCD) > types.Second { // 1s cooldown
				if !lonelyTimer.Stop() {
					<-lonelyTimer.C
				}
				lonelyTimer.Reset(timeout)
				resetCD = now
			}
			go handler.ServeWS(ws, msg)
		}
	}
}

// Watch performs basic socket i/o
func Watch(ws *T, handler Handler) {
	for {
		select {
		case buff := <-ws.send: // write to client
			if err := ws.Send(buff); err != nil {
				go drainChanMessage(ws.recv)
				ws.Close()
				return
			}
		case msg := <-ws.recv: // read from client
			if msg == nil {
				ws.Close()
				return
			}
			go handler.ServeWS(ws, msg)
		}
	}
}
