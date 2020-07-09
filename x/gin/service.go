package gin

import (
	"context"
	"sync"
	"time"

	"github.com/better-go/pkg/log"
	"github.com/better-go/pkg/os/signal"
	"github.com/better-go/pkg/x/gin/client"
	"github.com/better-go/pkg/x/gin/server"
)

// Service is an interface that wraps the lower level libraries
// within go-micro. Its a convenience method for building
// and initialising services.
type Service interface {
	// The service name
	Name() string
	// Init initialises options
	Init(...Option)
	// Options returns the current options
	Options() Options
	// Client is used to call services
	Client() client.Client
	// Server is for handling requests and events
	Server() server.Server
	// Run the service
	Run() error
	// The service implementation
	String() string
}

// NewService creates and returns a new Service based on the packages within.
func NewService(opts ...Option) Service {
	return newService(opts...)
}

//////////////////////////////////////////////////////////////////////////////////////////

type service struct {
	opts Options

	once sync.Once
}

func newService(opts ...Option) *service {
	s := new(service)

	// opt:
	o := newOptions(opts...)

	// set:
	s.opts = o

	return s
}

func (m *service) Name() string {
	return m.opts.Server.Options().Name
}

func (m *service) Init(opts ...Option) {
	// process options
	for _, o := range opts {
		o(&m.opts)
	}
}

func (m *service) Options() Options {
	return m.opts
}

// support graceful shutdown
func (m *service) Run() error {
	return signal.GracefulShutdown(
		m.opts.Context,
		func() error {
			//
			// do start:
			//
			return m.Start()
		},

		func() error {
			ctx, cancel := context.WithTimeout(context.Background(), 35*time.Second)

			//
			// do stop:
			//
			err := m.Stop(ctx)
			if err != nil {
				log.Errorf("http stop error: %v", err)
			}

			cancel()
			time.Sleep(time.Second)
			return err
		},
	)
}

func (m *service) RunTLS() error {
	// TODO: need fix
	return m.Run()
}

// start:
func (m *service) Start() error {
	// before:
	for _, fn := range m.opts.BeforeStart {
		if err := fn(); err != nil {
			return err
		}
	}

	// server:
	if err := m.Server().Start(); err != nil {
		return err
	}

	// after:
	for _, fn := range m.opts.AfterStart {
		if err := fn(); err != nil {
			return err
		}
	}

	return nil
}

func (m *service) Stop(ctx context.Context) error {
	var gerr error

	for _, fn := range m.opts.BeforeStop {
		if err := fn(); err != nil {
			gerr = err
		}
	}

	// server:
	if err := m.Server().Stop(ctx); err != nil {
		return err
	}

	for _, fn := range m.opts.AfterStop {
		if err := fn(); err != nil {
			gerr = err
		}
	}

	return gerr
}

func (m *service) Client() client.Client {
	return m.opts.Client
}

func (m *service) Server() server.Server {
	return m.opts.Server
}

func (m *service) String() string {
	return "micro"
}
