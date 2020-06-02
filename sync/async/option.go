package async

import (
	"context"
)

type options struct {
	worker int
	buffer int
}

// Option TaskDispatcher option
type Option func(*options)

// Worker specifies the worker of TaskDispatcher
func Worker(n int) Option {
	if n <= 0 {
		panic("fanout: worker should > 0")
	}
	return func(o *options) {
		o.worker = n
	}
}

// Buffer specifies the buffer of TaskDispatcher
func Buffer(n int) Option {
	if n <= 0 {
		panic("fanout: buffer should > 0")
	}
	return func(o *options) {
		o.buffer = n
	}
}

type item struct {
	taskFn func(c context.Context)
	ctx    context.Context
}
