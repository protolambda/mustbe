package be

import (
	"github.com/protolambda/mustbe/assertion"
)

// InMap asserts that the given key is present in the map.
func InMap[K comparable, V any](m map[K]V, key K) assertion.Assertion {
	return assertion.InMap[K, V]{Key: key, Map: m}
}
