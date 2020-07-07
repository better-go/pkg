package gin

import (
	"context"
	"net"
	"net/http"
	"net/http/pprof"
	"sync/atomic"
	"time"

	"github.com/better-go/pkg/log"
	"github.com/better-go/pkg/os/signal"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type Server struct {
	opts   Options      // server option
	server atomic.Value // store *http.Server

	r *gin.Engine // gin router
}

func NewServer(opts ...Option) *Server {
	s := new(Server)

	// opt:
	o := newOptions(opts...)

	// gin:
	r := gin.Default()

	// set:
	s.opts = o
	s.r = r

	return s
}

// support graceful shutdown
func (m *Server) Run() error {
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

func (m *Server) RunTLS() error {
	// TODO: need fix
	return m.Run()
}

// start:
func (m *Server) Start() error {
	// before:
	for _, fn := range m.opts.BeforeStart {
		if err := fn(); err != nil {
			return err
		}
	}

	////////////////////////////////////////////////////////////////////

	cfg := m.opts.Http
	l, err := net.Listen(cfg.Network, cfg.Addr)
	if err != nil {
		return errors.Wrapf(err, "listen tcp: %s", cfg.Addr)
	}
	log.Infof("http listen addr: %v", cfg)

	h := &http.Server{
		Addr:    m.opts.Http.Addr, // host
		Handler: m.r,              // register gin router

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

	////////////////////////////////////////////////////////////////////

	// after:
	for _, fn := range m.opts.AfterStart {
		if err := fn(); err != nil {
			return err
		}
	}

	return nil
}

func (m *Server) Stop(ctx context.Context) error {
	var gerr error

	for _, fn := range m.opts.BeforeStop {
		if err := fn(); err != nil {
			gerr = err
		}
	}

	////////////////////////////////////////////////////////////////////

	// stop:
	if err := m.shutdown(ctx); err != nil {
		log.Error("http shutdown error(%v)", err)
	}

	////////////////////////////////////////////////////////////////////

	for _, fn := range m.opts.AfterStop {
		if err := fn(); err != nil {
			gerr = err
		}
	}

	return gerr
}

func (m *Server) serve(s *http.Server, l net.Listener) error {
	s.Handler = m.r
	m.server.Store(s) // global cache

	// serve http:
	if err := s.Serve(l); err != nil {
		err = errors.WithStack(err)
		return err
	}
	return nil
}

func (m *Server) shutdown(ctx context.Context) error {
	s := m.Server()
	if s == nil {
		return errors.New("no server")
	}
	return errors.WithStack(s.Shutdown(ctx))
}

// Server is used to load stored http server.
func (m *Server) Server() *http.Server {
	s, ok := m.server.Load().(*http.Server)
	if !ok {
		return nil
	}
	return s
}

// go pref
func (m *Server) pref() {
	cfg := m.opts.HttpPref

	// do once:
	cfg.Once.Do(func() {
		// switch:
		if m.opts.HttpPref.On {
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
				if err := http.ListenAndServe(m.opts.HttpPref.Addr, nil); err != nil {
					panic(errors.Errorf("http pref listen %s: error(%v)", cfg.Addr, err))
				}
			}()
		}
	})
}
