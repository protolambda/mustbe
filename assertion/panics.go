package assertion

import (
	"context"
	"errors"
	"fmt"
	"runtime/debug"
)

func safeCall(fn func()) (panicked bool, out any, stack string) {
	panicked = true
	defer func() {
		out = recover()
		if panicked {
			stack = string(debug.Stack())
		}
	}()
	fn()
	panicked = false
	return
}

// Panic asserts that the given function panics.
type Panic struct {
	Fn func()
}

var _ Assertion = Panic{}

func (a Panic) String() string {
	return "isPanic"
}

func (a Panic) Check(ctx context.Context) error {
	panicked, _, _ := safeCall(a.Fn)
	if !panicked {
		return errors.New("expected function to panic")
	}
	return nil
}

// NoPanic asserts that the given function does not panic.
type NoPanic struct {
	Fn func()
}

var _ Assertion = NoPanic{}

func (a NoPanic) String() string {
	return "isNoPanic"
}

func (a NoPanic) Check(ctx context.Context) error {
	panicked, out, stack := safeCall(a.Fn)
	if panicked {
		return fmt.Errorf("expected function to not panic, but got %v\n%s", out, stack)
	}
	return nil
}
