package be

import "github.com/protolambda/mustbe/assertion"

// Panic asserts that the given function panics.
func Panic(fn func()) assertion.Assertion {
	return assertion.Panic{Fn: fn}
}

// NoPanic asserts that the given function does not panic.
func NoPanic(fn func()) assertion.Assertion {
	return assertion.NoPanic{Fn: fn}
}
