package assertion

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/protolambda/mustbe/internal/constraints"
)

// InDelta asserts that the given values are within the given delta. The delta bound is inclusive.
// The delta must be positive. If the delta is zero, the values must be equal.
type InDelta[V constraints.Number] struct {
	A     V
	B     V
	Delta V
}

var _ Assertion = InDelta[int]{}

func (s InDelta[V]) String() string {
	return fmt.Sprintf("isInDelta(%v, %v, delta: %v)", s.A, s.B, s.Delta)
}

func (s InDelta[V]) Check(ctx context.Context) error {
	a, b, delta := s.A, s.B, s.Delta
	var x any = V(0)
	switch x.(type) {
	case float32, float64:
		if x := float64(a); math.IsNaN(x) {
			return errors.New("a is NaN")
		}
		if x := float64(b); math.IsNaN(x) {
			return errors.New("b is NaN")
		}
		if x := float64(delta); math.IsNaN(x) {
			return errors.New("delta is NaN")
		}
	}
	if delta < 0 {
		return errors.New("delta is negative")
	}
	if a == b {
		return nil
	}
	if delta == 0 {
		return errors.New("a != b (delta = 0)")
	}
	if a < b {
		right := a + delta
		if right < a {
			// overflow, so b is within upper bound
			return nil
		}
		if b <= right {
			return nil
		}
		return fmt.Errorf("a < a + delta < b: %v < %v < %v", a, right, b)
	} else {
		// established: b < a
		left := a - delta
		if left > a {
			// underflow, so b is within lower bound
			return nil
		}
		if b >= left {
			return nil
		}
		return fmt.Errorf("b < a - delta < a: %v < %v < %v", b, left, a)
	}
}
