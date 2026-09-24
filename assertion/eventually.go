package assertion

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// Eventually runs a function repeatedly until it succeeds,
// the attempts are exhausted, or the context is done.
// The first attempt runs immediately, subsequent attempts are spaced by Tick.
// Tick and Attempts must be positive.
// If the function panics, the assertion fails immediately.
type Eventually struct {
	Fn       func(ctx context.Context) error
	Tick     time.Duration
	Attempts uint
}

var _ Assertion = Eventually{}

func (e Eventually) String() string {
	return fmt.Sprintf("Eventually(tick: %s, attempts: %d)", e.Tick, e.Attempts)
}

func (e Eventually) Check(ctx context.Context) error {
	if e.Tick <= 0 {
		return fmt.Errorf("invalid Eventually tick: %s, must be positive", e.Tick)
	}
	if e.Attempts == 0 {
		return errors.New("invalid Eventually attempts: 0, must be positive")
	}
	if err := ctx.Err(); err != nil {
		return err
	}
	ticker := time.NewTicker(e.Tick)
	defer ticker.Stop()
	for attempt := uint(1); ; attempt++ {
		err, panicked := e.run(ctx)
		if panicked {
			return fmt.Errorf("Eventually panicked (attempt %d/%d): %w", attempt, e.Attempts, err)
		}
		if err == nil {
			return nil
		}
		if attempt >= e.Attempts {
			return fmt.Errorf("attempts exhausted (%d), last error: %w", e.Attempts, err)
		}
		select {
		case <-ctx.Done():
			return fmt.Errorf("context done after %d/%d attempts (%w), last error: %w", attempt, e.Attempts, ctx.Err(), err)
		case <-ticker.C:
		}
	}
}

func (e Eventually) run(ctx context.Context) (out error, panicked bool) {
	panicked = true
	defer func() {
		if v := recover(); v != nil {
			if x, ok := v.(error); ok {
				out = x
			} else {
				out = fmt.Errorf("%v", v)
			}
		}
	}()
	out = e.Fn(ctx)
	panicked = false
	return
}
