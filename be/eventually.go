package be

import (
	"context"
	"time"

	"github.com/protolambda/mustbe/assertion"
)

// Eventually runs a function repeatedly until it succeeds,
// the attempts are exhausted, or the context is done.
// The first attempt runs immediately, subsequent attempts are spaced by tick.
// On failure, the last error of fn is included.
func Eventually(fn func(ctx context.Context) error, tick time.Duration, attempts uint) assertion.Assertion {
	return assertion.Eventually{Fn: fn, Tick: tick, Attempts: attempts}
}
