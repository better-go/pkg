package client

import (
	"context"
	"sync"
)

type Options struct {
	Network      string
	Addr         string
	Timeout      string
	ReadTimeout  string
	WriteTimeout string

	On   bool      // pref switch
	Once sync.Once // do once

	//
	//
	//
	Metadata map[string]string
	Name     string
	Id       string
	Version  string

	// Other options for implementations of the interface
	// can be stored in a context
	Context context.Context
}

func NewOptions(options ...Option) Options {
	opts := Options{
		Context: context.Background(),
	}

	for _, o := range options {
		o(&opts)
	}

	return opts
}
