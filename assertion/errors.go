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

// ErrContains asserts that the given error is non-nil and its message contains Sub.
type ErrContains struct {
	V   error
	Sub string
}

var _ Assertion = ErrContains{}

func (e ErrContains) String() string {
	return fmt.Sprintf("errContains(err: %v, sub: %q)", e.V, e.Sub)
}

func (e ErrContains) Check(ctx context.Context) error {
	if e.V == nil {
		return errors.New("expected non-nil error")
	}
	if !strings.Contains(e.V.Error(), e.Sub) {
		return fmt.Errorf("err %q does not contain substring %q", e.V, e.Sub)
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
