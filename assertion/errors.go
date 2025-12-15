package assertion

import (
	"context"
	"errors"
	"fmt"
)

// ErrorIs asserts that the given error is linked to the target error.
type ErrorIs struct {
	V      error
	Target error
}

var _ Assertion = ErrorIs{}

func (e ErrorIs) String() string {
	return fmt.Sprintf("isError(err: %v, target: %v)", e.V, e.Target)
}

func (e ErrorIs) Check(ctx context.Context) error {
	if !errors.Is(e.V, e.Target) {
		return fmt.Errorf("err %v is not err target %v", e.V, e.Target)
	}
	return nil
}

// NoError asserts that the given error is nil.
type NoError struct {
	V error
}

var _ Assertion = NoError{}

func (e NoError) String() string {
	return fmt.Sprintf("noError(%v)", e.V)
}

func (e NoError) Check(ctx context.Context) error {
	if e.V != nil {
		return fmt.Errorf("unexpected error: %v", e.V)
	}
	return nil
}
