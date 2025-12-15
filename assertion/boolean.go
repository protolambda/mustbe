package assertion

import (
	"context"
	"errors"
	"fmt"
)

// False asserts that the given value is false.
type False struct {
	V bool
}

var _ Assertion = False{}

func (a False) String() string {
	return fmt.Sprintf("isFalse(%v)", a.V)
}

func (a False) Check(ctx context.Context) error {
	if a.V {
		return errors.New("expected False")
	}
	return nil
}

// True asserts that the given value is true.
type True struct {
	V bool
}

var _ Assertion = True{}

func (a True) String() string {
	return fmt.Sprintf("isTrue(%v)", a.V)
}

func (a True) Check(ctx context.Context) error {
	if !a.V {
		return errors.New("expected True")
	}
	return nil
}
