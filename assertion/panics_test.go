package assertion_test

import (
	"context"
	"errors"
	"testing"

	"github.com/protolambda/mustbe/assertion"
)

func TestPanic(t *testing.T) {
	t.Run("Function panics with string", func(t *testing.T) {
		a := assertion.Panic{Fn: func() { panic("oops") }}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "isPanic")
	})
	t.Run("Function panics with error", func(t *testing.T) {
		a := assertion.Panic{Fn: func() { panic(errors.New("test error")) }}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "isPanic")
	})
	t.Run("Function panics with nil", func(t *testing.T) {
		var panicArg any // to work-around gopls panic lint warning
		a := assertion.Panic{Fn: func() { panic(panicArg) }}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "isPanic")
	})
	t.Run("Function does not panic", func(t *testing.T) {
		x := "hello"
		a := assertion.Panic{Fn: func() { x = "changed" }}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), "isPanic")
		expectString(t, x, "changed")
	})
}

func TestNoPanic(t *testing.T) {
	t.Run("Function panics with string", func(t *testing.T) {
		a := assertion.NoPanic{Fn: func() { panic("oops") }}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), "isNoPanic")
	})
	t.Run("Function panics with error", func(t *testing.T) {
		a := assertion.NoPanic{Fn: func() { panic(errors.New("test error")) }}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), "isNoPanic")
	})
	t.Run("Function panics with nil", func(t *testing.T) {
		var panicArg any // to work-around gopls panic lint warning
		a := assertion.NoPanic{Fn: func() { panic(panicArg) }}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), "isNoPanic")
	})
	t.Run("Function does not panic", func(t *testing.T) {
		x := "hello"
		a := assertion.NoPanic{Fn: func() { x = "changed" }}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "isNoPanic")
		expectString(t, x, "changed")
	})
}
