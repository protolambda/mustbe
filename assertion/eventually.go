package assertion

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// Eventually runs a function repeatedly until it succeeds or the context is done.
type Eventually struct {
	Fn       func(ctx context.Context) error
	Tick     time.Duration
	Attempts uint
}

var _ Assertion = (*Eventually)(nil)

func (e Eventually) String() string {
	return fmt.Sprintf("Eventually(tick: %s, attempts: %d)", e.Tick, e.Attempts)
}

func (e Eventually) Check(ctx context.Context) error {
	ticker := time.NewTicker(e.Tick)
	defer ticker.Stop()
	attempts := e.Attempts
	for {
		if attempts == 0 {
			return errors.New("attempts exhausted")
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			attempts--
			err, panicked := e.run(ctx)
			if panicked {
				return fmt.Errorf("Eventually panicked (%d attempts left): %w", attempts, err)
			} else if err != nil {
				continue
			} else {
				return nil
			}
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
