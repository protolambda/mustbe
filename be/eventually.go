package be

import (
	"context"
	"time"

	"github.com/protolambda/mustbe/assertion"
)

// Eventually runs a function repeatedly until it succeeds or the context is done.
func Eventually(fn func(ctx context.Context) error, tick time.Duration, attempts uint) assertion.Assertion {
	return assertion.Eventually{Fn: fn, Tick: tick, Attempts: attempts}
}
