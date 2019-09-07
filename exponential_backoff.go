package runner

import (
	"math"
	"time"
)

const base = 2.0

// ExponentialBackoffRetry is retry strategy which
// increase delay between each attempt
type ExponentialBackoffRetry struct {
	// retry counter
	attempt int

	// maximum retry allowed
	max int
}

// Next satisfy RetryStrategy.Next interface
func (r *ExponentialBackoffRetry) Next() time.Duration {
	r.attempt++

	// if attempt counter exceeded allowed
	if r.attempt > r.max {
		return Stop
	}

	return time.Duration(math.Pow(base, float64(r.attempt))) * time.Second
}

// Reset satisfy RetryStrategy.Reset interface
func (r *ExponentialBackoffRetry) Reset() {
	r.attempt = 0
}

// NewExponentialBackoffRetry create retry exponential
// backoff strategy
func NewExponentialBackoffRetry(max int) RetryStrategy {
	return &ExponentialBackoffRetry{max: max}
}
