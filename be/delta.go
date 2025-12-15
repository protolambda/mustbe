package be

import (
	"github.com/protolambda/mustbe/assertion"
	"github.com/protolambda/mustbe/internal/constraints"
)

// InDelta asserts that the given values are within the given delta.
// The delta must be positive. If the delta is zero, the values must be equal.
func InDelta[V constraints.Number](a, b, delta V) assertion.Assertion {
	return assertion.InDelta[V]{A: a, B: b, Delta: delta}
}
