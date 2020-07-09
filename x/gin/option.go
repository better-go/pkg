package gin

import (
	"context"
	"github.com/better-go/pkg/x/gin/client"
	"github.com/better-go/pkg/x/gin/server"
)

type Option func(*Options)

// Options for micro service
type Options struct {
	Server server.Server
	Client client.Client

	// Before and After funcs
	BeforeStart []func() error
	BeforeStop  []func() error
	AfterStart  []func() error
	AfterStop   []func() error

	// Other options for implementations of the interface
	// can be stored in a context
	Context context.Context

	Signal bool
}

func newOptions(opts ...Option) Options {
	opt := Options{
		Context: context.Background(),

		Server: server.DefaultServer,
		Client: client.DefaultClient,
	}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

// Server to be used for service
func Server(s server.Server) Option {
	return func(o *Options) {
		o.Server = s
	}
}

// Client to be used for service
func Client(c client.Client) Option {
	return func(o *Options) {
		o.Client = c
	}
}
