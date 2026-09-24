package assertion_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/protolambda/mustbe/assertion"
)

func TestEventually(t *testing.T) {
	t.Run("Succeeds on first attempt", func(t *testing.T) {
		a := assertion.Eventually{
			Fn:       func(ctx context.Context) error { return nil },
			Tick:     time.Millisecond,
			Attempts: 5,
		}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "Eventually(tick: 1ms, attempts: 5)")
	})
	t.Run("Succeeds after retries", func(t *testing.T) {
		attempts := 0
		a := assertion.Eventually{
			Fn: func(ctx context.Context) error {
				attempts++
				if attempts < 3 {
					return errors.New("not yet")
				}
				return nil
			},
			Tick:     time.Millisecond,
			Attempts: 5,
		}
		err := a.Check(context.Background())
		expectNoError(t, err)
		if attempts != 3 {
			t.Errorf("expected 3 attempts, got %d", attempts)
		}
	})
	t.Run("Exhausts attempts", func(t *testing.T) {
		attempts := 0
		a := assertion.Eventually{
			Fn: func(ctx context.Context) error {
				attempts++
				return errors.New("always fails")
			},
			Tick:     time.Millisecond,
			Attempts: 3,
		}
		err := a.Check(context.Background())
		expectError(t, err)
		if attempts != 3 {
			t.Errorf("expected 3 attempts, got %d", attempts)
		}
		expectString(t, err.Error(), "attempts exhausted (3), last error: always fails")
	})
	t.Run("Exhausted attempts wrap last error", func(t *testing.T) {
		sentinel := errors.New("sentinel")
		a := assertion.Eventually{
			Fn:       func(ctx context.Context) error { return sentinel },
			Tick:     time.Millisecond,
			Attempts: 2,
		}
		err := a.Check(context.Background())
		if !errors.Is(err, sentinel) {
			t.Errorf("expected wrapped sentinel, got %v", err)
		}
	})
	t.Run("First attempt runs immediately", func(t *testing.T) {
		a := assertion.Eventually{
			Fn:       func(ctx context.Context) error { return nil },
			Tick:     time.Hour,
			Attempts: 1,
		}
		start := time.Now()
		err := a.Check(context.Background())
		expectNoError(t, err)
		if d := time.Since(start); d > time.Minute {
			t.Errorf("expected immediate attempt, took %s", d)
		}
	})
	t.Run("Context done includes last error", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
		defer cancel()
		sentinel := errors.New("sentinel")
		a := assertion.Eventually{
			Fn:       func(ctx context.Context) error { return sentinel },
			Tick:     time.Hour,
			Attempts: 2,
		}
		err := a.Check(ctx)
		if !errors.Is(err, context.DeadlineExceeded) {
			t.Errorf("expected context.DeadlineExceeded, got %v", err)
		}
		if !errors.Is(err, sentinel) {
			t.Errorf("expected wrapped sentinel, got %v", err)
		}
	})
	t.Run("Non-positive tick", func(t *testing.T) {
		for _, tick := range []time.Duration{0, -time.Second} {
			a := assertion.Eventually{
				Fn:       func(ctx context.Context) error { return nil },
				Tick:     tick,
				Attempts: 1,
			}
			err := a.Check(context.Background())
			expectError(t, err)
		}
	})
	t.Run("Zero attempts", func(t *testing.T) {
		a := assertion.Eventually{
			Fn:       func(ctx context.Context) error { return nil },
			Tick:     time.Millisecond,
			Attempts: 0,
		}
		err := a.Check(context.Background())
		expectError(t, err)
	})
	t.Run("Context cancelled", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel() // cancel immediately
		a := assertion.Eventually{
			Fn:       func(ctx context.Context) error { return errors.New("fail") },
			Tick:     time.Millisecond,
			Attempts: 100,
		}
		err := a.Check(ctx)
		expectError(t, err)
		if !errors.Is(err, context.Canceled) {
			t.Errorf("expected context.Canceled, got %v", err)
		}
	})
	t.Run("Function panics", func(t *testing.T) {
		a := assertion.Eventually{
			Fn:       func(ctx context.Context) error { panic("oops") },
			Tick:     time.Millisecond,
			Attempts: 5,
		}
		err := a.Check(context.Background())
		expectError(t, err)
	})
	t.Run("Function panics with error", func(t *testing.T) {
		a := assertion.Eventually{
			Fn:       func(ctx context.Context) error { panic(errors.New("panic error")) },
			Tick:     time.Millisecond,
			Attempts: 5,
		}
		err := a.Check(context.Background())
		expectError(t, err)
	})
	t.Run("String format", func(t *testing.T) {
		a := assertion.Eventually{
			Fn:       func(ctx context.Context) error { return nil },
			Tick:     100 * time.Millisecond,
			Attempts: 10,
		}
		expectString(t, a.String(), "Eventually(tick: 100ms, attempts: 10)")
	})
}
