package retry

import (
	"time"
)

// WithRetry: support backOff, 自动重试, 支持回退策略
func WithSimpleRetry(retries int, taskFn func() error) error {
	// 默认策略:
	backOff := DefaultBackOffConfig

	// do task:
	if err := taskFn(); err != nil {
		if retries--; retries > 0 {
			time.Sleep(backOff.BackOff(retries))
			return WithSimpleRetry(retries, taskFn)
		}
		return err
	}
	return nil
}

// WithRetry: support backOff
func WithRetry(retries int, backOff BackOffConfig, taskFn func() error) error {
	// do task:
	if err := taskFn(); err != nil {
		if retries--; retries > 0 {
			time.Sleep(backOff.BackOff(retries))
			return WithRetry(retries, backOff, taskFn)
		}
		return err
	}
	return nil
}
