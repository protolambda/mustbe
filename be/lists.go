package be

import (
	"cmp"

	"github.com/protolambda/mustbe/assertion"
)

// Len asserts that the given list has the given length.
func Len[V any](list []V, expectedLen int) assertion.Assertion {
	return assertion.Len[V]{List: list, ExpectedLen: expectedLen}
}

// Empty asserts that the given list is empty.
func Empty[V any](list []V) assertion.Assertion {
	return assertion.Empty[V]{List: list}
}

// NonEmpty asserts that the given list is not empty.
func NonEmpty[V any](list []V) assertion.Assertion {
	return assertion.NonEmpty[V]{List: list}
}

// Ascending asserts that the given list is sorted in ascending order
// (each item is less or equal to the next item).
func Ascending[V cmp.Ordered](list []V) assertion.Assertion {
	return assertion.Ascending[V]{List: list}
}

// Descending asserts that the given list is sorted in descending order
// (each item is greater or equal to the next item).
func Descending[V cmp.Ordered](list []V) assertion.Assertion {
	return assertion.Descending[V]{List: list}
}

// InList asserts that the given item is contained in the list.
// See Contains, which is equivalent.
func InList[V comparable](list []V, item V) assertion.Assertion {
	return assertion.InList[V]{List: list, Item: item}
}

// Contains asserts that the given item is contained in the list.
// The argument order matches slices.Contains.
func Contains[V comparable](list []V, item V) assertion.Assertion {
	return assertion.Contains[V]{List: list, Item: item}
}
