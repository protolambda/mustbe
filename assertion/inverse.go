package assertion

import (
	"context"
	"errors"
	"fmt"
)

// Not asserts that the given assertion fails.
// Context cancellation or deadline errors are not treated as an inner failure,
// and are returned as-is instead.
// For better readability and error reporting, use the assertion-specific inverse assertions instead.
type Not struct {
	Inner Assertion
}

var _ Assertion = Not{}

func (a Not) String() string {
	return fmt.Sprintf("not(%v)", a.Inner.String())
}

func (a Not) Check(ctx context.Context) error {
	err := a.Inner.Check(ctx)
	if err == nil {
		return errors.New("expected inner assertion to fail")
	}
	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		return err
	}
	if ctxErr := ctx.Err(); ctxErr != nil {
		return fmt.Errorf("context done during inner assertion (%w): %w", ctxErr, err)
	}
	return nil
}
