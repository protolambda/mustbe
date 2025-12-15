package be

import "github.com/protolambda/mustbe/assertion"

// False asserts that the given value is false.
func False(v bool) assertion.Assertion {
	return assertion.False{V: v}
}

// True asserts that the given value is true.
func True(v bool) assertion.Assertion {
	return assertion.True{V: v}
}
