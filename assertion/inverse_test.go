package assertion_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/protolambda/mustbe/assertion"
)

func TestNot(t *testing.T) {
	t.Run("Not of Failed passes", func(t *testing.T) {
		a := assertion.Not{Inner: assertion.Failed{}}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "not(failed)")
	})
	t.Run("Not of Passed fails", func(t *testing.T) {
		a := assertion.Not{Inner: assertion.Passed{}}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), "not(passed)")
	})
	t.Run("Double negation of Failed fails", func(t *testing.T) {
		a := assertion.Not{Inner: assertion.Not{Inner: assertion.Failed{}}}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), "not(not(failed))")
	})
	t.Run("Double negation of Passed passes", func(t *testing.T) {
		a := assertion.Not{Inner: assertion.Not{Inner: assertion.Passed{}}}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "not(not(passed))")
	})
	t.Run("Not propagates context cancellation", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		a := assertion.Not{Inner: assertion.Eventually{
			Fn:       func(ctx context.Context) error { return nil },
			Tick:     time.Millisecond,
			Attempts: 3,
		}}
		err := a.Check(ctx)
		if !errors.Is(err, context.Canceled) {
			t.Errorf("expected context.Canceled, got %v", err)
		}
	})
	t.Run("Not fails if context is done after inner failure", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		a := assertion.Not{Inner: ctxFailAssertion{cancel: cancel}}
		err := a.Check(ctx)
		if !errors.Is(err, context.Canceled) {
			t.Errorf("expected context.Canceled, got %v", err)
		}
	})
}

// ctxFailAssertion cancels the context during the check,
// and returns an error that does not wrap the context error.
type ctxFailAssertion struct {
	cancel context.CancelFunc
}

func (a ctxFailAssertion) String() string {
	return "ctxFail"
}

func (a ctxFailAssertion) Check(ctx context.Context) error {
	a.cancel()
	return errors.New("operation aborted")
}
