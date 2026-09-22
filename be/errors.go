package be

import "github.com/protolambda/mustbe/assertion"

// ErrorIs asserts that the given error is linked to the target error.
func ErrorIs(v error, target error) assertion.Assertion {
	return assertion.ErrorIs{V: v, Target: target}
}

// ErrContains asserts that the given error is non-nil and its message contains sub.
func ErrContains(v error, sub string) assertion.Assertion {
	return assertion.ErrContains{V: v, Sub: sub}
}

// NoError asserts that the given error is nil.
func NoError(v error) assertion.Assertion {
	return assertion.NoError{V: v}
}
