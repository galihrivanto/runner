package runner

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// ErrGiveUp indicate no more retry needed
var ErrGiveUp = errors.New("retry stopped. give up")

// Operation is function to be executed via runner
// nil to indicate successful operation otherwise
// unsuccessfuly operation which may need to retried
type Operation func(context.Context) error

// SignalHandler handle OS signals
// caller client can intercept certain signal in order
// to execute certain operation like log rotate.
type SignalHandler interface {
	Handle(func(os.Signal))
}

// Stop define retry operation should be stopped
const Stop time.Duration = -1

// RetryStrategy define when should an operation
// should be retried unless received GiveUp error
type RetryStrategy interface {
	// Next return duration next operation
	// should be retried
	Next() time.Duration

	// Reset internal counter if applicable
	Reset()
}

// NoRetry indicate operation must be run once
// without retry
type NoRetry struct{}

// Next satisfy RetryStrategy.Next interface
func (r *NoRetry) Next() time.Duration {
	return Stop
}

// Reset satisfy RetryStrategy.Reset interface
func (r *NoRetry) Reset() {}

// Run execute operation excactly once without retry
func Run(ctx context.Context, operation Operation) SignalHandler {
	return RunWithRetry(ctx, operation, &NoRetry{})
}

type signalInterceptor struct {
	ctx context.Context
}

// Handle intercept OS signal and execute callback
func (r *signalInterceptor) Handle(fn func(s os.Signal)) {
	sink := make(chan os.Signal, 1)
	defer close(sink)

	// wait for signal
	signal.Notify(sink, signals...)

	// reset the watched signals
	defer signal.Ignore(signals...)

	for {
		select {
		case <-r.ctx.Done():
			return
		case sig := <-sink:
			if sig != syscall.SIGHUP {
				fn(sig)
				return
			}

			fn(sig)
		}
	}
}

// RunWithRetry execute operation with retry strategy.
// It could be NoRetry, Exponential Backoff or constant backoff
func RunWithRetry(ctx context.Context, operation Operation, retry RetryStrategy) SignalHandler {
	newCtx, cancel := context.WithCancel(ctx)

	go func() {
		defer cancel()

		for {
			// execute operation
			err := operation(ctx)
			if err == nil {
				// reset retry
				retry.Reset()
			}

			// if should ben stopped
			if err == nil || err == ErrGiveUp {
				cancel()
			}

			// wait until next iteration
			next := retry.Next()

			// if next retry is stop, then no more retry
			if next == Stop {
				cancel()
			}

			select {
			case <-newCtx.Done():
				return
			case <-time.After(next):
				break
			}
		}
	}()

	return &signalInterceptor{ctx: newCtx}
}
