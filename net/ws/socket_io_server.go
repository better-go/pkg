package ws

import (
	v2 "github.com/ambelovsky/gosf"
)

/*

ref:
	- https://github.com/googollee/go-socket.io
	- https://github.com/graarh/golang-socketio
	- https://github.com/ambelovsky/gosf

*/

type SocketIOMessageHandlerFunc func(client *v2.Client, request *v2.Request) *v2.Message

// server side:
type SocketIOServer struct {
	opt *SocketIOOption
}

func NewSocketIOServer(opt *SocketIOOption) *SocketIOServer {

	s := &SocketIOServer{
		opt: opt,
	}
	return s
}

func (m *SocketIOServer) RunV2(port int) {
	// endpoints:
	m.registerEndpoints()

	// http run:
	v2.Startup(map[string]interface{}{
		"port": port,
	})
}

func (m *SocketIOServer) registerEndpoints() {
	// register handler:
	for _, e := range m.opt.Endpoints {
		v2.Listen(e.Endpoint, e.Handler)
	}
}
