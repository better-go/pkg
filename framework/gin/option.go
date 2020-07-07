package gin

import (
	"context"
	"sync"
)

type Option func(*Options)

// Options for micro service
type Options struct {
	Http     ServerUnit // http server
	HttpPref ServerUnit // go pref

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

type ServerUnit struct {
	Network      string
	Addr         string
	Timeout      string
	ReadTimeout  string
	WriteTimeout string

	On   bool      // pref switch
	Once sync.Once // do once
}

func newOptions(opts ...Option) Options {
	opt := Options{
		Context: context.Background(),
		Http: ServerUnit{
			Addr:         "",
			Timeout:      "",
			ReadTimeout:  "",
			WriteTimeout: "",
			On:           false,
			Once:         sync.Once{},
		},
		HttpPref: ServerUnit{
			Addr:         "",
			Timeout:      "",
			ReadTimeout:  "",
			WriteTimeout: "",
			On:           true,
			Once:         sync.Once{},
		},

		Signal: true,
	}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}
