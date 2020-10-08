package server

import (
	"context"
	"net"
	"net/http"
	"net/http/pprof"
	"sync"
	"sync/atomic"

	"github.com/better-go/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

var (
	DefaultServer Server = newGinServer()
	// NewServer creates a new server
	NewServer func(...Option) Server = newGinServer
)

type ginServer struct {
	opts Options

	once   sync.Once
	server atomic.Value // store *http.Service
	r      *gin.Engine  // gin router

	exit chan chan error
	sync.RWMutex

	// graceful exit
	wg *sync.WaitGroup
}

func newGinServer(opts ...Option) Server {
	o := newOptions(opts...)

	return &ginServer{
		opts:    o,
		once:    sync.Once{},
		server:  atomic.Value{},
		r:       nil,
		exit:    nil,
		RWMutex: sync.RWMutex{},
		wg:      nil,
	}
}

func (m *ginServer) Options() Options {
	m.RLock()
	opts := m.opts
	m.RUnlock()
	return opts
}

func (m *ginServer) Init(opts ...Option) error {
	m.Lock()
	defer m.Unlock()

	for _, opt := range opts {
		opt(&m.opts)
	}

	return nil
}

func (m *ginServer) String() string {
	return "gin server"
}

// start:
func (m *ginServer) Start() error {

	////////////////////////////////////////////////////////////////////

	cfg := m.opts
	l, err := net.Listen(cfg.Network, cfg.Addr)
	if err != nil {
		return errors.Wrapf(err, "listen tcp: %s", cfg.Addr)
	}
	log.Infof("http listen addr: %v", cfg)

	h := &http.Server{
		Addr:    m.opts.Addr, // host
		Handler: m.r,         // register gin router

		TLSConfig:         nil,
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}

	// run:
	go func() {
		// run:
		if err := m.serve(h, l); err != nil {
			if errors.Cause(err) == http.ErrServerClosed {
				log.Info("http server closed")
				return
			}
			panic(errors.Wrapf(err, "http ListenServer(%+v, %+v)", h, l))
		}

	}()

	return nil
}

func (m *ginServer) Stop(ctx context.Context) error {
	// stop:
	err := m.shutdown(ctx)
	if err != nil {
		log.Error("http shutdown error(%v)", err)
	}
	return err
}

func (m *ginServer) serve(s *http.Server, l net.Listener) error {
	s.Handler = m.r
	m.server.Store(s) // global cache

	// serve http:
	if err := s.Serve(l); err != nil {
		err = errors.WithStack(err)
		return err
	}
	return nil
}

func (m *ginServer) shutdown(ctx context.Context) error {
	s := m.HttpServer()
	if s == nil {
		return errors.New("no server")
	}
	return errors.WithStack(s.Shutdown(ctx))
}

// Service is used to load stored http server.
func (m *ginServer) HttpServer() *http.Server {
	s, ok := m.server.Load().(*http.Server)
	if !ok {
		return nil
	}
	return s
}

// go pref
func (m *ginServer) pref() {
	cfg := m.opts

	// do once:
	cfg.Once.Do(func() {
		// switch:
		if m.opts.On {
			// register router:
			debug := m.r.Group("/debug/pprof")
			{
				debug.GET("/", pprofHandler(pprof.Index))
				debug.GET("/cmdline", pprofHandler(pprof.Cmdline))
				debug.GET("/profile", pprofHandler(pprof.Profile))
				debug.POST("/symbol", pprofHandler(pprof.Symbol))
				debug.GET("/symbol", pprofHandler(pprof.Symbol))
				debug.GET("/trace", pprofHandler(pprof.Trace))
				debug.GET("/allocs", pprofHandler(pprof.Handler("allocs").ServeHTTP))
				debug.GET("/block", pprofHandler(pprof.Handler("block").ServeHTTP))
				debug.GET("/goroutine", pprofHandler(pprof.Handler("goroutine").ServeHTTP))
				debug.GET("/heap", pprofHandler(pprof.Handler("heap").ServeHTTP))
				debug.GET("/mutex", pprofHandler(pprof.Handler("mutex").ServeHTTP))
				debug.GET("/threadcreate", pprofHandler(pprof.Handler("threadcreate").ServeHTTP))
			}

			// serve http:
			go func() {
				if err := http.ListenAndServe(m.opts.Addr, nil); err != nil {
					panic(errors.Errorf("http pref listen %s: error(%v)", cfg.Addr, err))
				}
			}()
		}
	})
}
