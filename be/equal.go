package be

import "github.com/protolambda/mustbe/assertion"

// Equal runs a shallow-equal assertion, using the Go == operator.
// Use DeepEqual for structural comparison.
func Equal[V comparable](expected V, got V) assertion.Assertion {
	return assertion.Equal[V]{Expected: expected, Got: got}
}

// NotEqual runs a shallow-not-equal assertion
func NotEqual[V comparable](unexpected V, got V) assertion.Assertion {
	return assertion.NotEqual[V]{Unexpected: unexpected, Got: got}
}

// DeepEqual runs a deep-equal assertion, using reflect.DeepEqual.
func DeepEqual[V any](expected V, got V) assertion.Assertion {
	return assertion.DeepEqual[V]{Expected: expected, Got: got}
}

// NotDeepEqual runs a deep-not-equal assertion
func NotDeepEqual[V any](unexpected V, got V) assertion.Assertion {
	return assertion.NotDeepEqual[V]{Unexpected: unexpected, Got: got}
}
