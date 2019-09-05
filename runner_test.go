package runner

import (
	"context"
	"os"
	"testing"
)

func TestNoRetry(t *testing.T) {
	runCount := 0

	op := func(ctx context.Context) error {
		runCount++

		return nil
	}

	Run(context.Background(), op).Handle(func(sig os.Signal) {
		// ignore
	})

	if runCount != 1 {
		t.Errorf("must be run only once, but run %d times", runCount)
	}
}
