package assertion

import (
	"bytes"
	"context"
	"fmt"
	"reflect"
)

// Equal runs a shallow-equal assertion, using the Go == operator.
// Pointers, channels and interfaces are compared by identity, not by the values they reference.
// Use DeepEqual for structural comparison.
type Equal[V comparable] struct {
	Expected V
	Got      V
}

var _ Assertion = Equal[int]{}

func (s Equal[V]) String() string {
	return fmt.Sprintf("isEqual(expected: %v, got: %v)", s.Expected, s.Got)
}

func (s Equal[V]) Check(ctx context.Context) error {
	if s.Expected != s.Got {
		return fmt.Errorf("not equal, expected: %v, got: %v", s.Expected, s.Got)
	}
	return nil
}

// NotEqual runs a shallow-not-equal assertion
type NotEqual[V comparable] struct {
	Unexpected V
	Got        V
}

var _ Assertion = NotEqual[int]{}

func (s NotEqual[V]) String() string {
	return fmt.Sprintf("isNotEqual(unexpected: %v, got: %v)", s.Unexpected, s.Got)
}

func (s NotEqual[V]) Check(ctx context.Context) error {
	if s.Unexpected == s.Got {
		return fmt.Errorf("equal, unexpected: %v, got: %v", s.Unexpected, s.Got)
	}
	return nil
}

// DeepEqual runs a deep-equal assertion, using reflect.DeepEqual
// (or bytes.Equal for byte slices, where nil and empty are equal).
type DeepEqual[V any] struct {
	Expected V
	Got      V
}

var _ Assertion = DeepEqual[int]{}

func (s DeepEqual[V]) String() string {
	return fmt.Sprintf("isDeepEqual(expected: %v, got: %v)", s.Expected, s.Got)
}

func (s DeepEqual[V]) Check(ctx context.Context) error {
	expected := s.Expected
	got := s.Got
	var x *V
	var xIface any = x
	if _, ok := xIface.(*[]byte); ok {
		var a, b any = expected, got
		if bytes.Equal(a.([]byte), b.([]byte)) {
			return nil
		}
		return fmt.Errorf("byte slices differ, expected: %x, got: %x", a, b)
	}
	if reflect.DeepEqual(expected, got) {
		return nil
	}
	return fmt.Errorf("not deep-equal, expected: %+v, got: %+v", expected, got)
}

// NotDeepEqual runs a deep-not-equal assertion
type NotDeepEqual[V any] struct {
	Unexpected V
	Got        V
}

var _ Assertion = NotDeepEqual[int]{}

func (s NotDeepEqual[V]) String() string {
	return fmt.Sprintf("isNotDeepEqual(unexpected: %v, got: %v)", s.Unexpected, s.Got)
}

func (s NotDeepEqual[V]) Check(ctx context.Context) error {
	unexpected := s.Unexpected
	got := s.Got
	var x *V
	var xIface any = x
	if _, ok := xIface.(*[]byte); ok {
		var a, b any = unexpected, got
		if bytes.Equal(a.([]byte), b.([]byte)) {
			return fmt.Errorf("byte slices equal, unexpected: %x, got: %x", a, b)
		}
		return nil
	}
	if reflect.DeepEqual(unexpected, got) {
		return fmt.Errorf("deep-equal, unexpected: %+v, got: %+v", unexpected, got)
	}
	return nil
}
