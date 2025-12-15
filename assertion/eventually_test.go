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
