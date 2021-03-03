package ws

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type WebSocketOption struct {
	Addr string
	Host string
	Port int

	//
	upgrade *websocket.Upgrader
}

func (m *WebSocketOption) Uri() string {
	if m.Addr != "" {
		return m.Addr
	}

	return fmt.Sprintf("%s:%d", m.Host, m.Port)
}
