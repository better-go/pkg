package main

import (
	v2 "github.com/ambelovsky/gosf"
	"github.com/better-go/pkg/log"
	"github.com/better-go/pkg/net/ws"
)

func main() {
	// TODO: need absolute path, change here! ! !
	homePage := "/Users/henry/Documents/iSpace/better-go/pkg/net/ws/example/socket_io/socket_io.html"

	log.Infof("socketIO html file path: %v", homePage)

	s := ws.NewSocketIOServer(&ws.SocketIOOption{
		Addr: "",
		Endpoints: []ws.SocketIOEndpoint{
			{
				Endpoint: "echo",
				Handler: func(client *v2.Client, request *v2.Request) *v2.Message {
					log.Infof("socketIO: receive message: %v", request.Message)
					return v2.NewSuccessMessage(request.Message.Text + " >>> echo from server")
				},
			},
		},
	})

	// run:
	s.RunV2(9999)
}
