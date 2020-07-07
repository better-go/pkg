package client

import (
	"sync/atomic"
)

var (
	// DefaultClient is a default client to use out of the box
	DefaultClient Client = newGinClient()
)

type ginClient struct {
	opts Options
	once atomic.Value
}

func newGinClient(opt ...Option) Client {
	opts := NewOptions(opt...)

	rc := &ginClient{
		opts: opts,
	}
	rc.once.Store(false)

	c := Client(rc)

	return c
}

func (r *ginClient) Init(opts ...Option) error {
	for _, o := range opts {
		o(&r.opts)
	}
	return nil
}

func (r *ginClient) Options() Options {
	return r.opts
}

func (r *ginClient) String() string {
	return "gin client"
}
