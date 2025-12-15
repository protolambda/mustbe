package assertion_test

import (
	"context"
	"testing"

	"github.com/protolambda/mustbe/assertion"
)

type Bar [3]int

type Foo struct {
	Bar Bar
}

var (
	fooZero    = Foo{}
	fooNonZero = Foo{Bar: Bar{1, 2, 3}}
)

func TestPositive(t *testing.T) {

	t.Run("Positive value", func(t *testing.T) {
		a := assertion.Positive[int]{V: 42}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "isPositive(got: 42)")
	})
	t.Run("Zero value", func(t *testing.T) {
		a := assertion.Positive[int]{V: 0}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), "isPositive(got: 0)")
	})
	t.Run("Negative value", func(t *testing.T) {
		a := assertion.Positive[int]{V: -5}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), "isPositive(got: -5)")
	})
	t.Run("Positive float", func(t *testing.T) {
		a := assertion.Positive[float64]{V: 3.14}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "isPositive(got: 3.14)")
	})
}

func TestNegative(t *testing.T) {
	t.Run("Negative value", func(t *testing.T) {
		a := assertion.Negative[int]{V: -42}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "isNegative(got: -42)")
	})
	t.Run("Zero value", func(t *testing.T) {
		a := assertion.Negative[int]{V: 0}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), "isNegative(got: 0)")
	})
	t.Run("Positive value", func(t *testing.T) {
		a := assertion.Negative[int]{V: 5}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), "isNegative(got: 5)")
	})
	t.Run("Negative float", func(t *testing.T) {
		a := assertion.Negative[float64]{V: -3.14}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "isNegative(got: -3.14)")
	})
}

func TestZero(t *testing.T) {
	t.Run("Zero int", func(t *testing.T) {
		a := assertion.Zero[int]{V: 0}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "isZero(got: 0)")
	})
	t.Run("Non-zero int", func(t *testing.T) {
		a := assertion.Zero[int]{V: 42}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), "isZero(got: 42)")
	})
	t.Run("Empty string", func(t *testing.T) {
		a := assertion.Zero[string]{V: ""}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "isZero(got: )")
	})
	t.Run("Non-empty string", func(t *testing.T) {
		a := assertion.Zero[string]{V: "hello"}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), "isZero(got: hello)")
	})
	t.Run("Zero struct", func(t *testing.T) {
		a := assertion.Zero[Foo]{V: fooZero}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "isZero(got: {[0 0 0]})")
	})
	t.Run("Non-zero struct", func(t *testing.T) {
		a := assertion.Zero[Foo]{V: fooNonZero}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), "isZero(got: {[1 2 3]})")
	})
}

func TestNotZero(t *testing.T) {
	t.Run("Non-zero int", func(t *testing.T) {
		a := assertion.NotZero[int]{V: 42}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "notZero(got: 42)")
	})
	t.Run("Zero int", func(t *testing.T) {
		a := assertion.NotZero[int]{V: 0}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), "notZero(got: 0)")
	})
	t.Run("Non-empty string", func(t *testing.T) {
		a := assertion.NotZero[string]{V: "hello"}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "notZero(got: hello)")
	})
	t.Run("Empty string", func(t *testing.T) {
		a := assertion.NotZero[string]{V: ""}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), "notZero(got: )")
	})
	t.Run("Non-zero struct", func(t *testing.T) {
		a := assertion.NotZero[Foo]{V: fooNonZero}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "notZero(got: {[1 2 3]})")
	})
	t.Run("Zero struct", func(t *testing.T) {
		a := assertion.NotZero[Foo]{V: fooZero}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), "notZero(got: {[0 0 0]})")
	})
}

func TestLess(t *testing.T) {
	t.Run("A less than B", func(t *testing.T) {
		a := assertion.Less[int]{A: 6, B: 7}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "less(a: 6, b: 7)")
	})
	t.Run("A equal to B", func(t *testing.T) {
		a := assertion.Less[int]{A: 6, B: 6}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), "less(a: 6, b: 6)")
	})
	t.Run("A greater than B", func(t *testing.T) {
		a := assertion.Less[int]{A: 7, B: 6}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), "less(a: 7, b: 6)")
	})
}

func TestLessOrEq(t *testing.T) {
	t.Run("A less than B", func(t *testing.T) {
		a := assertion.LessOrEq[int]{A: 6, B: 7}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "lessOrEq(a: 6, b: 7)")
	})
	t.Run("A equal to B", func(t *testing.T) {
		a := assertion.LessOrEq[int]{A: 6, B: 6}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "lessOrEq(a: 6, b: 6)")
	})
	t.Run("A greater than B", func(t *testing.T) {
		a := assertion.LessOrEq[int]{A: 7, B: 6}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), "lessOrEq(a: 7, b: 6)")
	})
}

func TestGreater(t *testing.T) {
	t.Run("A greater than B", func(t *testing.T) {
		a := assertion.Greater[int]{A: 7, B: 6}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "greater(a: 7, b: 6)")
	})
	t.Run("A equal to B", func(t *testing.T) {
		a := assertion.Greater[int]{A: 6, B: 6}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), "greater(a: 6, b: 6)")
	})
	t.Run("A less than B", func(t *testing.T) {
		a := assertion.Greater[int]{A: 6, B: 7}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), "greater(a: 6, b: 7)")
	})
}

func TestGreaterOrEq(t *testing.T) {
	t.Run("A greater than B", func(t *testing.T) {
		a := assertion.GreaterOrEq[int]{A: 7, B: 6}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "greaterOrEq(a: 7, b: 6)")
	})
	t.Run("A equal to B", func(t *testing.T) {
		a := assertion.GreaterOrEq[int]{A: 6, B: 6}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "greaterOrEq(a: 6, b: 6)")
	})
	t.Run("A less than B", func(t *testing.T) {
		a := assertion.GreaterOrEq[int]{A: 6, B: 7}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), "greaterOrEq(a: 6, b: 7)")
	})
}
