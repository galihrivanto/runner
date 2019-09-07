package runner

import (
	"time"
)

const defaultDelay = time.Second * 5

// ConstantBackoffRetry is retry strategy which
// use constant delay between each attempt
type ConstantBackoffRetry struct {
	// retry counter
	attempt int

	// maximum retry allowed
	max int

	// delay each attempt
	delay time.Duration
}

// Next satisfy RetryStrategy.Next interface
func (r *ConstantBackoffRetry) Next() time.Duration {
	r.attempt++

	// if attempt counter exceeded allowed
	if r.attempt > r.max {
		return Stop
	}

	return r.delay
}

// Reset satisfy RetryStrategy.Reset interface
func (r *ConstantBackoffRetry) Reset() {
	r.attempt = 0
}

// NewConstantBackoffRetry create retry exponential
// backoff strategy
func NewConstantBackoffRetry(max int, d ...time.Duration) RetryStrategy {
	delay := defaultDelay
	if len(d) > 0 && d[0] > 0 {
		delay = d[0]
	}

	return &ConstantBackoffRetry{max: max, delay: delay}
}
