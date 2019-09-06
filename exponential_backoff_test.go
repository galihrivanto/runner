package runner

import (
	"context"
	"os"
	"testing"
	"time"
)

var testCases = []time.Duration{
	time.Second * 2,
	time.Second * 4,
	time.Second * 8,
	time.Second * 16,
	time.Second * 32,
}

func TestExponentialBackoff(t *testing.T) {
	r := NewExponentialBackoffRetry(5)

	for _, tc := range testCases {
		next := r.Next()
		if next != tc {
			t.Errorf("Expected delay %v but recevied %v", tc, next)
		}
	}
}

func TestExponentialBackoffRetry(t *testing.T) {
	runCount := 0
	maxRun := 5

	op := func(ctx context.Context) error {
		runCount++

		return nil
	}

	RunWithRetry(context.Background(), op, NewExponentialBackoffRetry(maxRun)).Handle(func(sig os.Signal) {
		runCount++
	})

	if runCount != 1 {
		t.Errorf("expected run %d, but run %d times", maxRun, runCount)
	}
}
