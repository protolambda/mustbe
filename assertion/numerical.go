package assertion

import (
	"cmp"
	"context"
	"fmt"
)

// Positive asserts that V is greater than the zero value.
type Positive[V cmp.Ordered] struct {
	V V
}

var _ Assertion = Positive[int]{}

func (p Positive[V]) String() string {
	return fmt.Sprintf("isPositive(got: %v)", p.V)
}

func (p Positive[V]) Check(ctx context.Context) error {
	var zero V
	if p.V <= zero {
		return fmt.Errorf("expected positive value but got %v", p.V)
	}
	return nil
}

// Negative asserts that V is less than the zero value.
type Negative[V cmp.Ordered] struct {
	V V
}

var _ Assertion = Negative[int]{}

func (n Negative[V]) String() string {
	return fmt.Sprintf("isNegative(got: %v)", n.V)
}

func (n Negative[V]) Check(ctx context.Context) error {
	var zero V
	if n.V >= zero {
		return fmt.Errorf("expected negative value but got %v", n.V)
	}
	return nil
}

// Zero asserts that V is the zero value.
type Zero[V comparable] struct {
	V V
}

var _ Assertion = Zero[int]{}

func (z Zero[V]) String() string {
	return fmt.Sprintf("isZero(got: %v)", z.V)
}

func (z Zero[V]) Check(ctx context.Context) error {
	var zero V
	if zero != z.V {
		return fmt.Errorf("expected zero value but got %v", z.V)
	}
	return nil
}

// NotZero asserts that V is not the zero value.
type NotZero[V comparable] struct {
	V V
}

var _ Assertion = NotZero[int]{}

func (n NotZero[V]) String() string {
	return fmt.Sprintf("notZero(got: %v)", n.V)
}

func (n NotZero[V]) Check(ctx context.Context) error {
	var zero V
	if zero == n.V {
		return fmt.Errorf("expected non-zero value but got %v", n.V)
	}
	return nil
}

// Less asserts that A < B.
type Less[V cmp.Ordered] struct {
	A, B V
}

var _ Assertion = Less[int]{}

func (l Less[V]) String() string {
	return fmt.Sprintf("less(a: %v, b: %v)", l.A, l.B)
}

func (l Less[V]) Check(ctx context.Context) error {
	if l.A >= l.B {
		return fmt.Errorf("expected %v < %v", l.A, l.B)
	}
	return nil
}

// LessOrEq asserts that A <= B.
type LessOrEq[V cmp.Ordered] struct {
	A, B V
}

var _ Assertion = LessOrEq[int]{}

func (l LessOrEq[V]) String() string {
	return fmt.Sprintf("lessOrEq(a: %v, b: %v)", l.A, l.B)
}

func (l LessOrEq[V]) Check(ctx context.Context) error {
	if l.A > l.B {
		return fmt.Errorf("expected %v <= %v", l.A, l.B)
	}
	return nil
}

// Greater asserts that A > B.
type Greater[V cmp.Ordered] struct {
	A, B V
}

var _ Assertion = Greater[int]{}

func (g Greater[V]) String() string {
	return fmt.Sprintf("greater(a: %v, b: %v)", g.A, g.B)
}

func (g Greater[V]) Check(ctx context.Context) error {
	if g.A <= g.B {
		return fmt.Errorf("expected %v > %v", g.A, g.B)
	}
	return nil
}

// GreaterOrEq asserts that A >= B.
type GreaterOrEq[V cmp.Ordered] struct {
	A, B V
}

var _ Assertion = GreaterOrEq[int]{}

func (g GreaterOrEq[V]) String() string {
	return fmt.Sprintf("greaterOrEq(a: %v, b: %v)", g.A, g.B)
}

func (g GreaterOrEq[V]) Check(ctx context.Context) error {
	if g.A < g.B {
		return fmt.Errorf("expected %v >= %v", g.A, g.B)
	}
	return nil
}
