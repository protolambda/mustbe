package assertion

import (
	"context"
	"errors"
	"fmt"
)

// Not asserts that the given assertion fails.
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
	return nil
}
