package assertion

import (
	"context"
	"errors"
	"fmt"
	"strings"
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

// ErrorContains asserts that the given error is non-nil and its message contains Sub.
type ErrorContains struct {
	V   error
	Sub string
}

var _ Assertion = ErrorContains{}

func (e ErrorContains) String() string {
	return fmt.Sprintf("errorContains(err: %v, sub: %q)", e.V, e.Sub)
}

func (e ErrorContains) Check(ctx context.Context) error {
	if e.V == nil {
		return errors.New("expected non-nil error")
	}
	if !strings.Contains(e.V.Error(), e.Sub) {
		return fmt.Errorf("err %q does not contain substring %q", e.V, e.Sub)
	}
	return nil
}

// ErrContains asserts that the given error is non-nil and its message contains Sub.
//
// Deprecated: use ErrorContains instead.
type ErrContains = ErrorContains

// Error asserts that the given error is non-nil.
type Error struct {
	V error
}

var _ Assertion = Error{}

func (e Error) String() string {
	return fmt.Sprintf("hasError(%v)", e.V)
}

func (e Error) Check(ctx context.Context) error {
	if e.V == nil {
		return errors.New("expected an error but got nil")
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
