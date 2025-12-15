package assertion

import (
	"context"
	"fmt"
	"reflect"
)

// Nil asserts that the given value is nil, including
// interface values that contain nil pointers.
type Nil struct {
	V any
}

var _ Assertion = Nil{}

func (n Nil) String() string {
	return fmt.Sprintf("isNil(got: %v)", n.V)
}

func (n Nil) Check(ctx context.Context) error {
	if !isNil(n.V) {
		return fmt.Errorf("expected nil but got %v", n.V)
	}
	return nil
}

// NonNil asserts that the given value is not nil, including
// interface values that contain nil pointers.
type NonNil struct {
	V any
}

var _ Assertion = NonNil{}

func (n NonNil) String() string {
	return fmt.Sprintf("nonNil(got: %v)", n.V)
}

func (n NonNil) Check(ctx context.Context) error {
	if isNil(n.V) {
		return fmt.Errorf("expected non-nil value but got nil")
	}
	return nil
}

// NilPtr asserts that the given pointer is nil.
type NilPtr[V any] struct {
	P *V
}

var _ Assertion = NilPtr[int]{}

func (n NilPtr[V]) String() string {
	return fmt.Sprintf("nilPtr(got: %v)", n.P)
}

func (n NilPtr[V]) Check(ctx context.Context) error {
	if n.P != nil {
		return fmt.Errorf("expected nil pointer but got %v", n.P)
	}
	return nil
}

// NonNilPtr asserts that the given pointer is not nil.
type NonNilPtr[V any] struct {
	P *V
}

var _ Assertion = NonNilPtr[int]{}

func (n NonNilPtr[V]) String() string {
	return fmt.Sprintf("nonNilPtr(got: %v)", n.P)
}

func (n NonNilPtr[V]) Check(ctx context.Context) error {
	if n.P == nil {
		return fmt.Errorf("expected non-nil pointer but got nil")
	}
	return nil
}

// isNil checks if the given interface value is nil, including
// interface values that contain nil pointers.
func isNil(v any) bool {
	if v == nil {
		return true
	}
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Ptr, reflect.Map, reflect.Slice, reflect.Chan, reflect.Func, reflect.Interface:
		return rv.IsNil()
	}
	return false
}

// Same asserts that the pointers are pointing to the same thing.
type Same[V any] struct {
	Expected *V
	Got      *V
}

var _ Assertion = Same[int]{}

func (s Same[V]) String() string {
	return fmt.Sprintf("isSame(expected: %v, got: %v)", s.Expected, s.Got)
}

func (s Same[V]) Check(ctx context.Context) error {
	if s.Expected != s.Got {
		return fmt.Errorf("not the same pointer, expected: %v, got: %v", s.Expected, s.Got)
	}
	return nil
}
