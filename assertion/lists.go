package assertion

import (
	"cmp"
	"context"
	"fmt"
	"slices"
)

// Len asserts that the given list has the given length.
type Len[V any] struct {
	List        []V
	ExpectedLen int
}

var _ Assertion = Len[int]{}

func (l Len[V]) String() string {
	return fmt.Sprintf("isLen(expected: %d, got: %d)", l.ExpectedLen, len(l.List))
}

func (l Len[V]) Check(ctx context.Context) error {
	if len(l.List) != l.ExpectedLen {
		return fmt.Errorf("expected list length %d but got %d", l.ExpectedLen, len(l.List))
	}
	return nil
}

// Empty asserts that the given list is empty.
type Empty[V any] struct {
	List []V
}

var _ Assertion = Empty[int]{}

func (e Empty[V]) String() string {
	return fmt.Sprintf("isEmpty(got: %d)", len(e.List))
}

func (e Empty[V]) Check(ctx context.Context) error {
	if len(e.List) != 0 {
		return fmt.Errorf("expected empty list but got %d", len(e.List))
	}
	return nil
}

// NonEmpty asserts that the given list is not empty.
type NonEmpty[V any] struct {
	List []V
}

var _ Assertion = NonEmpty[int]{}

func (n NonEmpty[V]) String() string {
	return fmt.Sprintf("isNonEmpty(got: %d)", len(n.List))
}

func (n NonEmpty[V]) Check(ctx context.Context) error {
	if len(n.List) == 0 {
		return fmt.Errorf("expected non-empty list but got empty")
	}
	return nil
}

// Ascending asserts that the given list is sorted in ascending order
// (each item is less or equal to the next item).
type Ascending[V cmp.Ordered] struct {
	List []V
}

var _ Assertion = Ascending[int]{}

func (a Ascending[V]) String() string {
	return fmt.Sprintf("isAscending(%d items)", len(a.List))
}

func (a Ascending[V]) Check(ctx context.Context) error {
	for i := 1; i < len(a.List); i++ {
		if a.List[i-1] > a.List[i] {
			return fmt.Errorf("expected list to be ascending but got lower value at index %d", i)
		}
	}
	return nil
}

// Descending asserts that the given list is sorted in descending order
// (each item is greater or equal to the next item).
type Descending[V cmp.Ordered] struct {
	List []V
}

var _ Assertion = Descending[int]{}

func (d Descending[V]) String() string {
	return fmt.Sprintf("isDescending(%d items)", len(d.List))
}

func (d Descending[V]) Check(ctx context.Context) error {
	for i := 1; i < len(d.List); i++ {
		if d.List[i-1] < d.List[i] {
			return fmt.Errorf("expected list to be descending but got higher value at index %d", i)
		}
	}
	return nil
}

// InList asserts that V is contained in the list
type InList[V comparable] struct {
	List []V
	Item V
}

var _ Assertion = InList[int]{}

func (i InList[V]) String() string {
	return fmt.Sprintf("isInList(item: %v, list: %d items)", i.Item, len(i.List))
}

func (i InList[V]) Check(ctx context.Context) error {
	if slices.Contains(i.List, i.Item) {
		return nil
	}
	return fmt.Errorf("item %v not found in list", i.Item)
}
