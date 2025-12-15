package be

import "github.com/protolambda/mustbe/assertion"

// Nil asserts that the given value is nil, including
// interface values that contain nil pointers.
func Nil(v any) assertion.Assertion {
	return assertion.Nil{V: v}
}

// NonNil asserts that the given value is not nil, including
// interface values that contain nil pointers.
func NonNil(v any) assertion.Assertion {
	return assertion.NonNil{V: v}
}

// NilPtr asserts that the given pointer is nil.
func NilPtr[V any](v *V) assertion.Assertion {
	return assertion.NilPtr[V]{P: v}
}

// NonNilPtr asserts that the given pointer is not nil.
func NonNilPtr[V any](v *V) assertion.Assertion {
	return assertion.NonNilPtr[V]{P: v}
}

// Same asserts that the pointers are pointing to the same thing.
func Same[V any](expected *V, got *V) assertion.Assertion {
	return assertion.Same[V]{Expected: expected, Got: got}
}
