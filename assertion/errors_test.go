package assertion_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/protolambda/mustbe/assertion"
)

var (
	errSentinel = errors.New("sentinel error")
	errOther    = errors.New("other error")
)

func TestErrorIs(t *testing.T) {
	t.Run("Error matches target", func(t *testing.T) {
		a := assertion.ErrorIs{V: errSentinel, Target: errSentinel}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "isError(err: sentinel error, target: sentinel error)")
	})
	t.Run("Wrapped error matches target", func(t *testing.T) {
		wrapped := fmt.Errorf("wrapped: %w", errSentinel)
		a := assertion.ErrorIs{V: wrapped, Target: errSentinel}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "isError(err: wrapped: sentinel error, target: sentinel error)")
	})
	t.Run("Error does not match target", func(t *testing.T) {
		a := assertion.ErrorIs{V: errSentinel, Target: errOther}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), "isError(err: sentinel error, target: other error)")
	})
	t.Run("Nil error does not match target", func(t *testing.T) {
		a := assertion.ErrorIs{V: nil, Target: errSentinel}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), "isError(err: <nil>, target: sentinel error)")
	})
	t.Run("Nil error matches nil target", func(t *testing.T) {
		a := assertion.ErrorIs{V: nil, Target: nil}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "isError(err: <nil>, target: <nil>)")
	})
	t.Run("Deeply wrapped error matches target", func(t *testing.T) {
		wrapped := fmt.Errorf("level1: %w", fmt.Errorf("level2: %w", errSentinel))
		a := assertion.ErrorIs{V: wrapped, Target: errSentinel}
		err := a.Check(context.Background())
		expectNoError(t, err)
	})
}

func TestErrContains(t *testing.T) {
	t.Run("Error contains substring", func(t *testing.T) {
		a := assertion.ErrContains{V: errors.New("failed to read file"), Sub: "read"}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), `errContains(err: failed to read file, sub: "read")`)
	})
	t.Run("Error does not contain substring", func(t *testing.T) {
		a := assertion.ErrContains{V: errors.New("failed to read file"), Sub: "write"}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), `errContains(err: failed to read file, sub: "write")`)
	})
	t.Run("Nil error", func(t *testing.T) {
		a := assertion.ErrContains{V: nil, Sub: "failure"}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), `errContains(err: <nil>, sub: "failure")`)
	})
	t.Run("Empty substring", func(t *testing.T) {
		a := assertion.ErrContains{V: errSentinel, Sub: ""}
		err := a.Check(context.Background())
		expectNoError(t, err)
	})
	t.Run("Case sensitive", func(t *testing.T) {
		a := assertion.ErrContains{V: errors.New("Failed"), Sub: "failed"}
		err := a.Check(context.Background())
		expectError(t, err)
	})
}

func TestNoError(t *testing.T) {
	t.Run("Nil error", func(t *testing.T) {
		a := assertion.NoError{V: nil}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "noError(<nil>)")
	})
	t.Run("Non-nil error", func(t *testing.T) {
		a := assertion.NoError{V: errSentinel}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), "noError(sentinel error)")
	})
	t.Run("Wrapped error", func(t *testing.T) {
		wrapped := fmt.Errorf("wrapped: %w", errSentinel)
		a := assertion.NoError{V: wrapped}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), "noError(wrapped: sentinel error)")
	})
}
