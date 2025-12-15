package mustbe

import (
	"context"
	"fmt"

	"github.com/protolambda/mustbe/assertion"
)

type T interface {
	Error(args ...any)
	FailNow()
	Context() context.Context
	Helper()
}

// Must is a shorthand to check an assertion immediately.
// For multiple assertion checks, wrap T into MT, and run `t.Must(...)` more seamlessly.
func Must(t T, c assertion.Assertion) {
	mustf(t, c, "")
}

func mustf(t T, c assertion.Assertion, msg string, args ...any) {
	defer func() {
		e := recover()
		if e != nil {
			t.Error("panic in assertion", e)
			t.FailNow()
		}
	}()
	ctx := t.Context()
	err := c.Check(ctx)
	if err != nil {
		t.Helper()
		if msg != "" {
			err = fmt.Errorf("%w: %s", err, fmt.Sprintf(msg, args...))
		}
		t.Error("assertion failed:", err)
		t.FailNow()
	}
}

// MT is the T interface with the addition of Must and Mustf
type MT interface {
	T
	Must(c assertion.Assertion)
	Mustf(c assertion.Assertion, msg string, args ...any)
}

type mustT struct {
	T
}

var _ MT = mustT{}

func (m mustT) Must(c assertion.Assertion) {
	mustf(m.T, c, "")
}

func (m mustT) Mustf(c assertion.Assertion, msg string, args ...any) {
	mustf(m.T, c, msg, args...)
}

func WrapT(t T) MT {
	return mustT{T: t}
}
