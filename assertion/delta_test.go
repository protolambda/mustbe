package assertion_test

import (
	"context"
	"math"
	"testing"

	"github.com/protolambda/mustbe/assertion"
)

func TestInDelta(t *testing.T) {
	t.Run("Integers within delta", func(t *testing.T) {
		a := assertion.InDelta[int]{A: 10, B: 12, Delta: 5}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "isInDelta(10, 12, delta: 5)")
	})
	t.Run("Integers outside delta", func(t *testing.T) {
		a := assertion.InDelta[int]{A: 10, B: 20, Delta: 5}
		err := a.Check(context.Background())
		expectError(t, err)
	})
	t.Run("Integers exactly at delta boundary", func(t *testing.T) {
		a := assertion.InDelta[int]{A: 10, B: 15, Delta: 5}
		err := a.Check(context.Background())
		expectNoError(t, err)
	})
	t.Run("Equal integers zero delta", func(t *testing.T) {
		a := assertion.InDelta[int]{A: 42, B: 42, Delta: 0}
		err := a.Check(context.Background())
		expectNoError(t, err)
	})
	t.Run("Unequal integers zero delta", func(t *testing.T) {
		a := assertion.InDelta[int]{A: 42, B: 43, Delta: 0}
		err := a.Check(context.Background())
		expectError(t, err)
	})
	t.Run("Negative delta", func(t *testing.T) {
		a := assertion.InDelta[int]{A: 10, B: 12, Delta: -5}
		err := a.Check(context.Background())
		expectError(t, err)
	})
	t.Run("A greater than B within delta", func(t *testing.T) {
		a := assertion.InDelta[int]{A: 15, B: 10, Delta: 5}
		err := a.Check(context.Background())
		expectNoError(t, err)
	})
	t.Run("A greater than B outside delta", func(t *testing.T) {
		a := assertion.InDelta[int]{A: 20, B: 10, Delta: 5}
		err := a.Check(context.Background())
		expectError(t, err)
	})
	t.Run("Floats within delta", func(t *testing.T) {
		a := assertion.InDelta[float64]{A: 1.0, B: 1.05, Delta: 0.1}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "isInDelta(1, 1.05, delta: 0.1)")
	})
	t.Run("Floats outside delta", func(t *testing.T) {
		a := assertion.InDelta[float64]{A: 1.0, B: 1.5, Delta: 0.1}
		err := a.Check(context.Background())
		expectError(t, err)
	})
	t.Run("Float32 within delta", func(t *testing.T) {
		a := assertion.InDelta[float32]{A: 1.0, B: 1.05, Delta: 0.1}
		err := a.Check(context.Background())
		expectNoError(t, err)
	})
	t.Run("A is NaN", func(t *testing.T) {
		a := assertion.InDelta[float64]{A: math.NaN(), B: 1.0, Delta: 0.1}
		err := a.Check(context.Background())
		expectError(t, err)
	})
	t.Run("B is NaN", func(t *testing.T) {
		a := assertion.InDelta[float64]{A: 1.0, B: math.NaN(), Delta: 0.1}
		err := a.Check(context.Background())
		expectError(t, err)
	})
	t.Run("Delta is NaN", func(t *testing.T) {
		a := assertion.InDelta[float64]{A: 1.0, B: 1.05, Delta: math.NaN()}
		err := a.Check(context.Background())
		expectError(t, err)
	})
	t.Run("Unsigned integers within delta", func(t *testing.T) {
		a := assertion.InDelta[uint]{A: 10, B: 12, Delta: 5}
		err := a.Check(context.Background())
		expectNoError(t, err)
	})
	t.Run("Unsigned integers outside delta", func(t *testing.T) {
		a := assertion.InDelta[uint]{A: 10, B: 20, Delta: 5}
		err := a.Check(context.Background())
		expectError(t, err)
	})
	t.Run("Integer overflow protection", func(t *testing.T) {
		// When a + delta overflows, b should still be within bounds
		a := assertion.InDelta[int8]{A: 120, B: 127, Delta: 100}
		err := a.Check(context.Background())
		expectNoError(t, err)
	})
	t.Run("Integer underflow protection", func(t *testing.T) {
		// When a - delta underflows, b should still be within bounds
		a := assertion.InDelta[int8]{A: -120, B: -128, Delta: 100}
		err := a.Check(context.Background())
		expectNoError(t, err)
	})
	t.Run("Integer overflow avoided", func(t *testing.T) {
		// Large delta, overflow if added to B instead of A
		a := assertion.InDelta[int8]{A: -40, B: 20, Delta: 127}
		err := a.Check(context.Background())
		expectNoError(t, err)
	})
}
