package ws

import (
	"net/http"

	"github.com/better-go/pkg/log"
	"github.com/gorilla/websocket"
)

/*

ref:
	- https://github.com/gorilla/websocket
		- usage: https://github.com/gorilla/websocket/blob/master/examples/echo/README.md
	- https://github.com/nhooyr/websocket
	- https://github.com/gobwas/ws

docs:
	- https://tonybai.com/2019/09/28/how-to-build-websockets-in-go/
	- https://cloud.tencent.com/developer/article/1605145
*/

// 接收消息+处理+生成待发送消息:
type MessageHandler func(receivedMessage []byte) (responseMessage []byte, err error)

// server side:
type WebSocketServer struct {
	//
	upgrade *websocket.Upgrader

	conn *websocket.Conn
}

func NewWebSocketServer() *WebSocketServer {
	s := &WebSocketServer{
		upgrade: new(websocket.Upgrader), // use default options
	}
	return s
}

func (m *WebSocketServer) ServeHTTP(w http.ResponseWriter, r *http.Request, handler MessageHandler) {
	conn, err := m.upgrade.Upgrade(w, r, nil)
	if err != nil {
		log.Errorf("websocket upgrade error: %v", err)
		return
	}
	defer conn.Close()

	m.conn = conn

	// handler ws request:
	m.dispatch(handler)
}

func (m *WebSocketServer) dispatch(handler MessageHandler) {
	// loop:
	for {
		// 1.  Read message from browser
		msgType, msg, err := m.conn.ReadMessage()
		if err != nil {
			log.Errorf("ws server: read message error: %v", err)
			return
		}
		log.Infof("ws server: read message from [%s]: msg=%s, type=%v", m.conn.RemoteAddr(), string(msg), msgType)

		// 2. parse message, and gen sent response
		sentMsg, err := handler(msg)
		if err != nil {
			log.Errorf("ws server: handler message error: req=%v, resp=%v, err=%v", msg, sentMsg, err)
			return
		}
		if sentMsg == nil {
			return
		}

		// 2. do send: Write message back to browser
		if err = m.conn.WriteMessage(msgType, sentMsg); err != nil {
			log.Errorf("ws server: send message back error: req=%v, type=%v, resp=%v, err=%v", msg, msgType, sentMsg, err)
			return
		}
	}
}
