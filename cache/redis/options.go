package redis

import (
	"time"

	sdk1 "github.com/go-redis/redis/v8"
)

const (
	defaultAddr         = ":6379"
	defaultDialTimeout  = 10 * time.Second
	defaultReadTimeout  = 30 * time.Second
	defaultWriteTimeout = 30 * time.Second
	defaultPoolSize     = 10
	defaultPoolTimeout  = 30 * time.Second
)

type Options struct {
	//
	Name  string // redis alias name, for trace
	Proto string //
	Addr  string
	Auth  string

	//
	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	SlowLog      time.Duration
	PoolSize     int

	//
	// Active number of items allocated by the pool at a given time.
	// When zero, there is no limit on the number of items in the pool.
	Active int
	// Idle number of idle items in the pool.
	Idle int
	// Close items after remaining item for this duration. If the value
	// is zero, then item items are not closed. Applications should set
	// the timeout to a value less than the server's timeout.
	IdleTimeout time.Duration
	// If WaitTimeout is set and the pool is at the Active limit, then Get() waits WatiTimeout
	// until a item to be returned to the pool before returning.
	WaitTimeout time.Duration
	// If WaitTimeout is not set, then Wait effects.
	// if Wait is set true, then wait until ctx timeout, or default flase and return directly.
	Wait bool
}

// fn:
type OptionFunc func(*Options)

// init default conf, update conf:
func (m *Options) Init(opts ...OptionFunc) {
	if m.Addr == "" {
		m.Addr = defaultAddr
	}
	if m.DialTimeout <= 0 {
		m.DialTimeout = defaultDialTimeout
	}
	if m.ReadTimeout <= 0 {
		m.ReadTimeout = defaultReadTimeout
	}
	if m.WriteTimeout <= 0 {
		m.WriteTimeout = defaultWriteTimeout
	}
	if m.PoolSize <= 0 {
		m.PoolSize = defaultPoolSize
	}

	// update:
	for _, fn := range opts {
		fn(m)
	}
	return
}

func (m *Options) ToOption1() *sdk1.Options {
	m.Init()

	return &sdk1.Options{
		Addr:         m.Addr,
		Password:     m.Auth,
		DB:           0,
		DialTimeout:  m.DialTimeout,
		ReadTimeout:  m.ReadTimeout,
		WriteTimeout: m.WriteTimeout,
		PoolSize:     m.PoolSize,
	}
}

func (m *Options) ToOption2() {
	m.Init()

}
