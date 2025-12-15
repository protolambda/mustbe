package assertion_test

import (
	"context"
	"testing"

	"github.com/protolambda/mustbe/assertion"
)

func TestOfType(t *testing.T) {
	t.Run("Matching type", func(t *testing.T) {
		a := assertion.OfType[int]{V: 42}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "isOfType(v: 42, type: int)")
	})
	t.Run("Non-matching type", func(t *testing.T) {
		a := assertion.OfType[int]{V: "hello"}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), "isOfType(v: hello, type: int)")
	})
	t.Run("Nil value", func(t *testing.T) {
		a := assertion.OfType[int]{V: nil}
		err := a.Check(context.Background())
		expectError(t, err)
	})
}

type foobarInterface interface {
	Foobar()
}

type foobarImplementer struct{}

func (foobarImplementer) Foobar() {}

type foobarNonImplementer struct{}

func TestImplemented(t *testing.T) {
	t.Run("Implements interface", func(t *testing.T) {
		a := assertion.Implemented[foobarInterface]{V: foobarImplementer{}}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "isImplemented(v: {}, interface: assertion_test.foobarInterface)")
	})
	t.Run("Does not implement interface", func(t *testing.T) {
		a := assertion.Implemented[foobarInterface]{V: foobarNonImplementer{}}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), "isImplemented(v: {}, interface: assertion_test.foobarInterface)")
	})
	t.Run("Nil value", func(t *testing.T) {
		a := assertion.Implemented[foobarInterface]{V: nil}
		err := a.Check(context.Background())
		expectError(t, err)
	})
}
