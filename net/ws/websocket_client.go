package ws

import "github.com/gorilla/websocket"

// client side:
type WebSocketClient struct {
	//
	upgrade *websocket.Upgrader

	conn *websocket.Conn
}

func NewWebSocketClient() *WebSocketClient {
	s := &WebSocketClient{
		upgrade: new(websocket.Upgrader), // use default options
	}
	return s
}
