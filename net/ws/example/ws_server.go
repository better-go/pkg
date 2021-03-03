package main

import (
	"github.com/better-go/pkg/log"
	"github.com/better-go/pkg/net/ws"
	"net/http"
)

func main() {
	// TODO: need absolute path, change here! ! !
	homePage := "/Users/henry/Documents/iSpace/better-go/pkg/net/ws/example/ws.html"

	log.Infof("ws html file path: %v", homePage)

	// ws echo endpoint:
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		s := ws.NewWebSocketServer()

		// loop do:
		s.ServeHTTP(w, r, func(receivedMessage []byte) (responseMessage []byte, err error) {
			log.Infof("ws server: receive message: %v", string(receivedMessage))
			resp := string(receivedMessage) + "### echo from server"
			responseMessage = []byte(resp)
			return responseMessage, nil
		})

	})

	// home page:
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, homePage)
	})

	// http server:
	http.ListenAndServe(":8080", nil)
}
