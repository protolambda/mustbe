package be

import (
	"cmp"

	"github.com/protolambda/mustbe/assertion"
)

// Positive asserts that v is greater than the zero value.
func Positive[V cmp.Ordered](v V) assertion.Assertion {
	return assertion.Positive[V]{V: v}
}

// Negative asserts that v is less than the zero value.
func Negative[V cmp.Ordered](v V) assertion.Assertion {
	return assertion.Negative[V]{V: v}
}

// Zero asserts that v is the zero value.
func Zero[V comparable](v V) assertion.Assertion {
	return assertion.Zero[V]{V: v}
}

// NotZero asserts that v is not the zero value.
func NotZero[V comparable](v V) assertion.Assertion {
	return assertion.NotZero[V]{V: v}
}

// Less asserts that a < b.
func Less[V cmp.Ordered](a, b V) assertion.Assertion {
	return assertion.Less[V]{A: a, B: b}
}

// LessOrEq asserts that a <= b.
func LessOrEq[V cmp.Ordered](a, b V) assertion.Assertion {
	return assertion.LessOrEq[V]{A: a, B: b}
}

// Greater asserts that a > b.
func Greater[V cmp.Ordered](a, b V) assertion.Assertion {
	return assertion.Greater[V]{A: a, B: b}
}

// GreaterOrEq asserts that a >= b.
func GreaterOrEq[V cmp.Ordered](a, b V) assertion.Assertion {
	return assertion.GreaterOrEq[V]{A: a, B: b}
}
