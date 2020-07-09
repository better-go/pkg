package server

import (
	"context"
)

// Server is a simple micro server abstraction
type Server interface {
	// Initialise options
	Init(...Option) error
	// Retrieve the options
	Options() Options

	// Start the server
	Start() error
	// Stop the server
	Stop(ctx context.Context) error
	// Server implementation
	String() string
}

type Option func(*Options)

type HandlerOption func(*HandlerOptions)

type HandlerOptions struct {
	Internal bool
	Metadata map[string]map[string]string
}

type SubscriberOption func(*SubscriberOptions)

type SubscriberOptions struct {
	// AutoAck defaults to true. When a handler returns
	// with a nil error the message is acked.
	AutoAck  bool
	Queue    string
	Internal bool
	Context  context.Context
}

// Handler interface represents a request handler. It's generated
// by passing any type of public concrete object with endpoints into server.NewHandler.
// Most will pass in a struct.
//
// Example:
//
//      type Greeter struct {}
//
//      func (g *Greeter) Hello(context, request, response) error {
//              return nil
//      }
//
type Handler interface {
	Name() string
	Handler() interface{}
	Options() HandlerOptions
}

// Subscriber interface represents a subscription to a given topic using
// a specific subscriber function or object with endpoints. It mirrors
// the handler in its behaviour.
type Subscriber interface {
	Topic() string
	Subscriber() interface{}
	Options() SubscriberOptions
}
