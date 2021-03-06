package runner

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"
)

var exponentialTestCases = []time.Duration{
	time.Second * 2,
	time.Second * 4,
	time.Second * 8,
	time.Second * 16,
	time.Second * 32,
}

func TestExponentialBackoff(t *testing.T) {
	r := NewExponentialBackoffRetry(5)

	for _, tc := range exponentialTestCases {
		next := r.Next()
		if next != tc {
			t.Errorf("Expected delay %v but recevied %v", tc, next)
			t.FailNow()
		}
	}
}

func TestExponentialBackoffRetry(t *testing.T) {
	runCount := 0
	maxRun := 3

	op := func(ctx context.Context) error {
		runCount++

		return fmt.Errorf("dummy error to retry")
	}

	// total run = first op + retry count
	RunWithRetry(context.Background(), op, NewExponentialBackoffRetry(maxRun-1)).
		Handle(func(sig os.Signal) {
			// ignore
		})

	if runCount != maxRun {
		t.Errorf("expected run %d, but run %d times", maxRun, runCount)
	}
}
