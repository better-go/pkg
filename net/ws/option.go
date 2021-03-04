package ws

import (
	"encoding/json"
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

// ws message format:
type Message struct {
	data []byte
}

func (m *Message) FromJson(from string) (err error) {
	// from -> data
	m.data, err = json.Marshal(from)
	return
}

func (m *Message) ToJson() (to map[string]interface{}, err error) {
	// data -> json string
	err = json.Unmarshal(m.data, &to)
	return
}

// protobuf:
func (m *Message) FromProtoBuf(from string) (err error) {
	// from -> data
	m.data, err = json.Marshal(from)
	return
}

func (m *Message) ToProtoBuf() (to map[string]interface{}, err error) {
	// data -> json string
	err = json.Unmarshal(m.data, &to)
	return
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////

type SocketIOOption struct {
	Addr string

	Endpoints []SocketIOEndpoint
}

type SocketIOEndpoint struct {
	Endpoint string
	Handler  SocketIOMessageHandlerFunc
}
