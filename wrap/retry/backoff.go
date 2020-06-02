package retry

import (
	"math/rand"
	"time"
)

var DefaultBackOffConfig = BackOffConfig{
	MaxDelay:  120 * time.Second,
	BaseDelay: 1.0 * time.Second,
	Factor:    1.6,
	Jitter:    0.2,
}

// BackOff defines the methodology for backing off after a call failure.
type BackOff interface {
	// BackOff returns the amount of time to wait before the next retry given
	// the number of consecutive failures.
	BackOff(retries int) time.Duration
}

// BackOffConfig defines the parameters for the default BackOff strategy.
type BackOffConfig struct {
	// MaxDelay is the upper bound of BackOff delay.
	MaxDelay time.Duration

	// baseDelay is the amount of time to wait before retrying after the first
	// failure.
	BaseDelay time.Duration

	// factor is applied to the BackOff after each retry.
	Factor float64

	// jitter provides a range to randomize BackOff delays.
	Jitter float64
}

// BackOff returns the amount of time to wait before the next retry given
// the number of consecutive failures.
func (bc *BackOffConfig) BackOff(retries int) time.Duration {
	if retries == 0 {
		return bc.BaseDelay
	}
	backOff, max := float64(bc.BaseDelay), float64(bc.MaxDelay)
	for backOff < max && retries > 0 {
		backOff *= bc.Factor
		retries--
	}
	if backOff > max {
		backOff = max
	}
	// Randomize BackOff delays so that if a cluster of requests start at
	// the same time, they won't operate in lockstep.
	backOff *= 1 + bc.Jitter*(rand.Float64()*2-1)
	if backOff < 0 {
		return 0
	}
	return time.Duration(backOff)
}
