package runner

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"
)

var constantTestCases = []time.Duration{
	time.Second * 2,
	time.Second * 2,
	time.Second * 2,
	time.Second * 2,
	time.Second * 2,
}

func TestConstantBackoff(t *testing.T) {
	r := NewConstantBackoffRetry(5, time.Second*2)

	for _, tc := range constantTestCases {
		next := r.Next()
		if next != tc {
			t.Errorf("Expected delay %v but recevied %v", tc, next)
			t.FailNow()
		}
	}
}

func TestConstantBackoffRetry(t *testing.T) {
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
