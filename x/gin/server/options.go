package server

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

// Init initialises the default server with options passed in
func Init(opt ...Option) {
	if DefaultServer == nil {
		DefaultServer = newGinServer(opt...)
	}
	DefaultServer.Init(opt...)
}

// String returns name of Server implementation
func String() string {
	return DefaultServer.String()
}

func newOptions(opt ...Option) Options {
	opts := Options{
		Network:      "",
		Addr:         "",
		Timeout:      "",
		ReadTimeout:  "",
		WriteTimeout: "",
		On:           true,
		Once:         sync.Once{},
		Metadata:     map[string]string{},
		Name:         "",
		Id:           "",
		Version:      "",
		Context:      nil,
	}

	for _, o := range opt {
		o(&opts)
	}

	return opts
}
