package be

import "github.com/protolambda/mustbe/assertion"

// OfType asserts that the given value is of the given type.
// The reflect type of T is compared to the reflect type of the given value.
func OfType[T any](v any) assertion.Assertion {
	return assertion.OfType[T]{V: v}
}

// Implemented asserts that the given value implements the given interface.
func Implemented[T any](v any) assertion.Assertion {
	return assertion.Implemented[T]{V: v}
}
