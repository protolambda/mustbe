package mustbe

import (
	"context"
	"runtime/debug"

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
//
// On failure, only the error returned by the assertion Check is reported:
// assertions include the relevant values in their errors,
// and the String representation is reserved for display and composition.
func Must(t T, c assertion.Assertion) {
	t.Helper()
	mustf(t, c, "")
}

func mustf(t T, c assertion.Assertion, msg string, args ...any) {
	t.Helper()
	if msg != "" {
		c = assertion.Annotated{Inner: c, Msg: msg, Args: args}
	}
	err, panicked, panicValue, stack := check(t.Context(), c)
	if panicked {
		// Reported outside of the deferred recover,
		// so the failure location is the Must call, not the panic site.
		t.Error("panic in assertion", panicValue, "\n"+stack)
		t.FailNow()
		return
	}
	if err != nil {
		t.Error("assertion failed:", err)
		t.FailNow()
	}
}

// check runs the assertion, and recovers from any panic in it.
func check(ctx context.Context, c assertion.Assertion) (err error, panicked bool, panicValue any, stack string) {
	panicked = true
	defer func() {
		if panicked {
			panicValue = recover()
			stack = string(debug.Stack())
		}
	}()
	err = c.Check(ctx)
	panicked = false
	return
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
	m.T.Helper()
	mustf(m.T, c, "")
}

func (m mustT) Mustf(c assertion.Assertion, msg string, args ...any) {
	m.T.Helper()
	mustf(m.T, c, msg, args...)
}

func WrapT(t T) MT {
	return mustT{T: t}
}
