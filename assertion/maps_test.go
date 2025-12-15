package assertion_test

import (
	"context"
	"testing"

	"github.com/protolambda/mustbe/assertion"
)

func TestInMap(t *testing.T) {
	t.Run("Key in map", func(t *testing.T) {
		a := assertion.InMap[string, int]{Map: map[string]int{"a": 1, "b": 2, "c": 3}, Key: "b"}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "isInMap(key: b, map: 3 entries)")
	})
	t.Run("Key not in map", func(t *testing.T) {
		a := assertion.InMap[string, int]{Map: map[string]int{"a": 1, "b": 2, "c": 3}, Key: "z"}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), "isInMap(key: z, map: 3 entries)")
	})
	t.Run("Empty map", func(t *testing.T) {
		a := assertion.InMap[string, int]{Map: map[string]int{}, Key: "a"}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), "isInMap(key: a, map: 0 entries)")
	})
	t.Run("Nil map", func(t *testing.T) {
		a := assertion.InMap[string, int]{Map: nil, Key: "a"}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), "isInMap(key: a, map: 0 entries)")
	})
	t.Run("Integer keys", func(t *testing.T) {
		a := assertion.InMap[int, string]{Map: map[int]string{1: "one", 2: "two", 3: "three"}, Key: 2}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "isInMap(key: 2, map: 3 entries)")
	})
	t.Run("Integer key not found", func(t *testing.T) {
		a := assertion.InMap[int, string]{Map: map[int]string{1: "one", 2: "two", 3: "three"}, Key: 99}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), "isInMap(key: 99, map: 3 entries)")
	})
	t.Run("Key with empty string value", func(t *testing.T) {
		a := assertion.InMap[string, int]{Map: map[string]int{"": 0}, Key: ""}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "isInMap(key: , map: 1 entries)")
	})
}
