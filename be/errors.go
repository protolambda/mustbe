package be

import "github.com/protolambda/mustbe/assertion"

// ErrorIs asserts that the given error is linked to the target error.
func ErrorIs(v error, target error) assertion.Assertion {
	return assertion.ErrorIs{V: v, Target: target}
}

// ErrorContains asserts that the given error is non-nil and its message contains sub.
func ErrorContains(v error, sub string) assertion.Assertion {
	return assertion.ErrorContains{V: v, Sub: sub}
}

// ErrContains asserts that the given error is non-nil and its message contains sub.
//
// Deprecated: use ErrorContains instead.
func ErrContains(v error, sub string) assertion.Assertion {
	return ErrorContains(v, sub)
}

// NoError asserts that the given error is nil.
func NoError(v error) assertion.Assertion {
	return assertion.NoError{V: v}
}

// Error asserts that the given error is non-nil.
func Error(v error) assertion.Assertion {
	return assertion.Error{V: v}
}
