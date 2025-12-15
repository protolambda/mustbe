package be

import "github.com/protolambda/mustbe/assertion"

// ErrorIs asserts that the given error is linked to the target error.
func ErrorIs(v error, target error) assertion.Assertion {
	return assertion.ErrorIs{V: v, Target: target}
}

// NoError asserts that the given error is nil.
func NoError(v error) assertion.Assertion {
	return assertion.NoError{V: v}
}
