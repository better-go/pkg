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

// ws header auth: by api-key
type AuthApiKeyValidateFunc func(apiKey string) bool

// server side:
type WebSocketServer struct {
	upgrade *websocket.Upgrader //
	conn    *websocket.Conn     //

	msgHandler    MessageHandler         // msg parse
	authValidator AuthApiKeyValidateFunc // auth check
}

func NewWebSocketServer(msgHandler MessageHandler, authValidator AuthApiKeyValidateFunc) *WebSocketServer {
	s := &WebSocketServer{
		upgrade:       new(websocket.Upgrader), // use default options
		msgHandler:    msgHandler,              // 只能在此处注册: 消息的处理(接收+处理+返回响应)
		authValidator: authValidator,           // 鉴权方式
	}
	return s
}

// 鉴权: http header key
func (m *WebSocketServer) DispatchWithAuth(middleware ...func(http.Handler) http.Handler) (h http.Handler) {
	// auth check:
	h = m.authMiddleware(http.HandlerFunc(m.Dispatch))

	// register others:
	for _, m := range middleware {
		h = m(h)
	}
	return h
}

// 处理消息:
func (m *WebSocketServer) Dispatch(w http.ResponseWriter, r *http.Request) {
	conn, err := m.upgrade.Upgrade(w, r, nil)
	if err != nil {
		log.Errorf("websocket upgrade error: %v", err)
		return
	}
	defer conn.Close()

	m.conn = conn

	// handler ws request:
	m.dispatch()
}

func (m *WebSocketServer) dispatch() {
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
		sentMsg, err := m.msgHandler(msg)
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

func (m *WebSocketServer) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// http header auth check:
		if apiKey := req.Header.Get("X-Api-Key"); !m.authValidator(apiKey) {
			log.Infof("ws auth failed, bad auth api key: %s", apiKey)
			rw.WriteHeader(http.StatusForbidden)
			return
		}

		//
		// do next step:
		//
		next.ServeHTTP(rw, req)
	})
}

func Middleware(h http.Handler, middleware ...func(http.Handler) http.Handler) http.Handler {
	for _, m := range middleware {
		h = m(h)
	}
	return h
}

// demo check:
func demoApiKeyCheck(apiKey string) bool {
	//
	// TODO: query server, validate api-key
	//
	key := "test_api_key_server_cache"
	return apiKey == key
}
