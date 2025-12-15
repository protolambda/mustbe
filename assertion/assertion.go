package assertion

import "context"

// Assertion assertions are deferred as objects,
// to allow for Go-generics type-safe assertions,
// and to make user-provided assertions fit in easily.
type Assertion interface {
	// String presents a readable version of the condition.
	String() string
	// Check runs the assertion.
	// This returns an error if the context is canceled or the assertion fails.
	Check(ctx context.Context) error
}
