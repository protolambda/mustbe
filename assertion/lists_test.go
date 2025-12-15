package assertion_test

import (
	"context"
	"testing"

	"github.com/protolambda/mustbe/assertion"
)

func TestLen(t *testing.T) {
	t.Run("Correct lenth", func(t *testing.T) {
		a := assertion.Len[int]{List: []int{1, 2, 3}, ExpectedLen: 3}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "isLen(expected: 3, got: 3)")
	})
	t.Run("Wrong lenth", func(t *testing.T) {
		a := assertion.Len[int]{List: []int{1, 2}, ExpectedLen: 5}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), "isLen(expected: 5, got: 2)")
	})
	t.Run("Empty list with zero expected", func(t *testing.T) {
		a := assertion.Len[int]{List: []int{}, ExpectedLen: 0}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "isLen(expected: 0, got: 0)")
	})
	t.Run("Empty list with non-zero expected", func(t *testing.T) {
		a := assertion.Len[string]{List: []string{}, ExpectedLen: 2}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), "isLen(expected: 2, got: 0)")
	})
	t.Run("Nil list", func(t *testing.T) {
		a := assertion.Len[int]{List: nil, ExpectedLen: 0}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "isLen(expected: 0, got: 0)")
	})
}

func TestEmpty(t *testing.T) {
	t.Run("Empty list", func(t *testing.T) {
		a := assertion.Empty[int]{List: []int{}}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "isEmpty(got: 0)")
	})
	t.Run("Nil list", func(t *testing.T) {
		a := assertion.Empty[int]{List: nil}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "isEmpty(got: 0)")
	})
	t.Run("Non-empty list", func(t *testing.T) {
		a := assertion.Empty[int]{List: []int{1, 2, 3}}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), "isEmpty(got: 3)")
	})
	t.Run("Single element", func(t *testing.T) {
		a := assertion.Empty[string]{List: []string{"hello"}}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), "isEmpty(got: 1)")
	})
}

func TestNonEmpty(t *testing.T) {
	t.Run("Non-empty list", func(t *testing.T) {
		a := assertion.NonEmpty[int]{List: []int{1, 2, 3}}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "isNonEmpty(got: 3)")
	})
	t.Run("Single element", func(t *testing.T) {
		a := assertion.NonEmpty[string]{List: []string{"hello"}}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "isNonEmpty(got: 1)")
	})
	t.Run("Empty list", func(t *testing.T) {
		a := assertion.NonEmpty[int]{List: []int{}}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), "isNonEmpty(got: 0)")
	})
	t.Run("Nil list", func(t *testing.T) {
		a := assertion.NonEmpty[int]{List: nil}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), "isNonEmpty(got: 0)")
	})
}

func TestAscending(t *testing.T) {
	t.Run("Ascending integers", func(t *testing.T) {
		a := assertion.Ascending[int]{List: []int{1, 2, 3, 4, 5}}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "isAscending(5 items)")
	})
	t.Run("Ascending with equal values", func(t *testing.T) {
		a := assertion.Ascending[int]{List: []int{1, 2, 2, 3, 3, 3}}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "isAscending(6 items)")
	})
	t.Run("Not ascending", func(t *testing.T) {
		a := assertion.Ascending[int]{List: []int{1, 3, 2}}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), "isAscending(3 items)")
	})
	t.Run("Descending order", func(t *testing.T) {
		a := assertion.Ascending[int]{List: []int{5, 4, 3, 2, 1}}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), "isAscending(5 items)")
	})
	t.Run("Empty list", func(t *testing.T) {
		a := assertion.Ascending[int]{List: []int{}}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "isAscending(0 items)")
	})
	t.Run("Single element", func(t *testing.T) {
		a := assertion.Ascending[int]{List: []int{42}}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "isAscending(1 items)")
	})
	t.Run("Ascending floats", func(t *testing.T) {
		a := assertion.Ascending[float64]{List: []float64{1.1, 2.2, 3.3}}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "isAscending(3 items)")
	})
	t.Run("Ascending strings", func(t *testing.T) {
		a := assertion.Ascending[string]{List: []string{"apple", "banana", "cherry"}}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "isAscending(3 items)")
	})
}

func TestDescending(t *testing.T) {
	t.Run("Descending integers", func(t *testing.T) {
		a := assertion.Descending[int]{List: []int{5, 4, 3, 2, 1}}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "isDescending(5 items)")
	})
	t.Run("Descending with equal values", func(t *testing.T) {
		a := assertion.Descending[int]{List: []int{5, 5, 4, 3, 3}}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "isDescending(5 items)")
	})
	t.Run("Not descending", func(t *testing.T) {
		a := assertion.Descending[int]{List: []int{3, 1, 2}}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), "isDescending(3 items)")
	})
	t.Run("Ascending order", func(t *testing.T) {
		a := assertion.Descending[int]{List: []int{1, 2, 3, 4, 5}}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), "isDescending(5 items)")
	})
	t.Run("Empty list", func(t *testing.T) {
		a := assertion.Descending[int]{List: []int{}}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "isDescending(0 items)")
	})
	t.Run("Single element", func(t *testing.T) {
		a := assertion.Descending[int]{List: []int{42}}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "isDescending(1 items)")
	})
	t.Run("Descending floats", func(t *testing.T) {
		a := assertion.Descending[float64]{List: []float64{3.3, 2.2, 1.1}}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "isDescending(3 items)")
	})
	t.Run("Descending strings", func(t *testing.T) {
		a := assertion.Descending[string]{List: []string{"cherry", "banana", "apple"}}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "isDescending(3 items)")
	})
}

func TestInList(t *testing.T) {
	t.Run("Item in list", func(t *testing.T) {
		a := assertion.InList[int]{List: []int{1, 2, 3, 4, 5}, Item: 3}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "isInList(item: 3, list: 5 items)")
	})
	t.Run("Item not in list", func(t *testing.T) {
		a := assertion.InList[int]{List: []int{1, 2, 3, 4, 5}, Item: 99}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), "isInList(item: 99, list: 5 items)")
	})
	t.Run("Item at beginning", func(t *testing.T) {
		a := assertion.InList[string]{List: []string{"apple", "banana", "cherry"}, Item: "apple"}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "isInList(item: apple, list: 3 items)")
	})
	t.Run("Item at end", func(t *testing.T) {
		a := assertion.InList[string]{List: []string{"apple", "banana", "cherry"}, Item: "cherry"}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "isInList(item: cherry, list: 3 items)")
	})
	t.Run("Empty list", func(t *testing.T) {
		a := assertion.InList[int]{List: []int{}, Item: 42}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), "isInList(item: 42, list: 0 items)")
	})
	t.Run("Nil list", func(t *testing.T) {
		a := assertion.InList[int]{List: nil, Item: 42}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), "isInList(item: 42, list: 0 items)")
	})
	t.Run("Single element found", func(t *testing.T) {
		a := assertion.InList[int]{List: []int{42}, Item: 42}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "isInList(item: 42, list: 1 items)")
	})
	t.Run("Single element not found", func(t *testing.T) {
		a := assertion.InList[int]{List: []int{99}, Item: 42}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), "isInList(item: 42, list: 1 items)")
	})
}
