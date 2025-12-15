package be

import "github.com/protolambda/mustbe/assertion"

// Not asserts that the given assertion fails.
// For better readability and error reporting, use the assertion-specific inverse assertions instead.
func Not(v assertion.Assertion) assertion.Assertion {
	return assertion.Not{Inner: v}
}
