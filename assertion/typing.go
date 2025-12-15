package assertion

import (
	"context"
	"fmt"
	"reflect"
)

// OfType asserts that the given value is of the given type.
// The reflect type of T is compared to the reflect type of the given value.
type OfType[T any] struct {
	V any
}

func (s OfType[T]) String() string {
	return fmt.Sprintf("isOfType(v: %v, type: %s)", s.V, reflect.TypeFor[T]())
}

func (s OfType[T]) Check(ctx context.Context) error {
	expected := reflect.TypeFor[T]()
	got := reflect.TypeOf(s.V)
	if !reflect.DeepEqual(expected, got) {
		return fmt.Errorf("expected type %s but got %s", expected, got)
	}
	return nil
}

// Implemented asserts that the given value implements the given interface.
type Implemented[T any] struct {
	V any
}

func (s Implemented[T]) String() string {
	return fmt.Sprintf("isImplemented(v: %v, interface: %s)", s.V, reflect.TypeFor[T]())
}

func (s Implemented[T]) Check(ctx context.Context) error {
	_, ok := s.V.(T)
	if !ok {
		return fmt.Errorf("%v (type %T) does not implement interface type %s", s.V, s.V, reflect.TypeFor[T]())
	}
	return nil
}
