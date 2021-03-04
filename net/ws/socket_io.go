package ws

import "github.com/gorilla/websocket"

/*

ref:
	- https://github.com/googollee/go-socket.io
	- https://github.com/graarh/golang-socketio
	- https://github.com/ambelovsky/gosf

*/

// server side:
type SocketIOServer struct {
	//
	upgrade *websocket.Upgrader

	conn *websocket.Conn
}

func NewSocketIOServer() *SocketIOServer {
	s := &SocketIOServer{
		upgrade: new(websocket.Upgrader), // use default options
	}
	return s
}



