package assertion_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/protolambda/mustbe/assertion"
)

func TestAnnotated(t *testing.T) {
	t.Run("Annotated passing assertion", func(t *testing.T) {
		a := assertion.Annotated{
			Inner: assertion.Passed{},
			Msg:   "test message",
		}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "test message: passed")
	})
	t.Run("Annotated failing assertion", func(t *testing.T) {
		a := assertion.Annotated{
			Inner: assertion.Failed{},
			Msg:   "test message",
		}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), "test message: failed")
	})
	t.Run("Annotated with format args", func(t *testing.T) {
		a := assertion.Annotated{
			Inner: assertion.Passed{},
			Msg:   "value is %d and name is %s",
			Args:  []any{42, "foo"},
		}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "value is 42 and name is foo: passed")
	})
	t.Run("Annotated with empty message", func(t *testing.T) {
		a := assertion.Annotated{
			Inner: assertion.Passed{},
			Msg:   "",
		}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), ": passed")
	})
	t.Run("Nested annotated", func(t *testing.T) {
		a := assertion.Annotated{
			Inner: assertion.Annotated{
				Inner: assertion.Passed{},
				Msg:   "inner",
			},
			Msg: "outer",
		}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "outer: inner: passed")
	})
	t.Run("Annotated with real assertion", func(t *testing.T) {
		a := assertion.Annotated{
			Inner: assertion.Equal[int]{Expected: 42, Got: 42},
			Msg:   "checking answer",
		}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "checking answer: isEqual(expected: 42, got: 42)")
	})
	t.Run("Annotated failure includes message", func(t *testing.T) {
		a := assertion.Annotated{
			Inner: assertion.Equal[int]{Expected: 1, Got: 2},
			Msg:   "value %d",
			Args:  []any{7},
		}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, err.Error(), "not equal, expected: 1, got: 2: value 7")
	})
	t.Run("Annotated failure wraps inner error", func(t *testing.T) {
		a := assertion.Annotated{
			Inner: assertion.Eventually{Fn: func(ctx context.Context) error { return nil }, Tick: time.Millisecond, Attempts: 1},
			Msg:   "msg",
		}
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		err := a.Check(ctx)
		if !errors.Is(err, context.Canceled) {
			t.Errorf("expected wrapped context.Canceled, got %v", err)
		}
	})
}
