package assertion_test

import (
	"context"
	"testing"

	"github.com/protolambda/mustbe/assertion"
)

func TestEqual(t *testing.T) {
	t.Run("Equal integers", func(t *testing.T) {
		a := assertion.Equal[int]{Expected: 42, Got: 42}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "isEqual(expected: 42, got: 42)")
	})
	t.Run("Unequal integers", func(t *testing.T) {
		a := assertion.Equal[int]{Expected: 42, Got: 99}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), "isEqual(expected: 42, got: 99)")
	})
	t.Run("Equal strings", func(t *testing.T) {
		a := assertion.Equal[string]{Expected: "hello", Got: "hello"}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "isEqual(expected: hello, got: hello)")
	})
	t.Run("Unequal strings", func(t *testing.T) {
		a := assertion.Equal[string]{Expected: "hello", Got: "world"}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), "isEqual(expected: hello, got: world)")
	})
	t.Run("Equal zero values", func(t *testing.T) {
		a := assertion.Equal[int]{Expected: 0, Got: 0}
		err := a.Check(context.Background())
		expectNoError(t, err)
	})
	t.Run("Equal booleans", func(t *testing.T) {
		a := assertion.Equal[bool]{Expected: true, Got: true}
		err := a.Check(context.Background())
		expectNoError(t, err)
	})
	t.Run("Unequal booleans", func(t *testing.T) {
		a := assertion.Equal[bool]{Expected: true, Got: false}
		err := a.Check(context.Background())
		expectError(t, err)
	})
}

type testStruct struct {
	Name  string
	Value int
}

func TestDeepEqual(t *testing.T) {
	t.Run("Equal structs", func(t *testing.T) {
		a := assertion.DeepEqual[testStruct]{
			Expected: testStruct{Name: "foo", Value: 42},
			Got:      testStruct{Name: "foo", Value: 42},
		}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "isDeepEqual(expected: {foo 42}, got: {foo 42})")
	})
	t.Run("Unequal structs", func(t *testing.T) {
		a := assertion.DeepEqual[testStruct]{
			Expected: testStruct{Name: "foo", Value: 42},
			Got:      testStruct{Name: "bar", Value: 99},
		}
		err := a.Check(context.Background())
		expectError(t, err)
	})
	t.Run("Equal slices", func(t *testing.T) {
		a := assertion.DeepEqual[[]int]{
			Expected: []int{1, 2, 3},
			Got:      []int{1, 2, 3},
		}
		err := a.Check(context.Background())
		expectNoError(t, err)
	})
	t.Run("Unequal slices", func(t *testing.T) {
		a := assertion.DeepEqual[[]int]{
			Expected: []int{1, 2, 3},
			Got:      []int{1, 2, 4},
		}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, err.Error(), "not deep-equal, expected: [1 2 3], got: [1 2 4]")
	})
	t.Run("Unequal structs show fields", func(t *testing.T) {
		type item struct {
			Name string
			N    int
		}
		a := assertion.DeepEqual[item]{Expected: item{Name: "a", N: 1}, Got: item{Name: "a", N: 2}}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, err.Error(), "not deep-equal, expected: {Name:a N:1}, got: {Name:a N:2}")
	})
	t.Run("Equal byte slices", func(t *testing.T) {
		a := assertion.DeepEqual[[]byte]{
			Expected: []byte{0x01, 0x02, 0x03},
			Got:      []byte{0x01, 0x02, 0x03},
		}
		err := a.Check(context.Background())
		expectNoError(t, err)
	})
	t.Run("Unequal byte slices", func(t *testing.T) {
		a := assertion.DeepEqual[[]byte]{
			Expected: []byte{0x01, 0x02, 0x03},
			Got:      []byte{0x01, 0x02, 0x04},
		}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, err.Error(), "byte slices differ, expected: 010203, got: 010204")
	})
	t.Run("Equal maps", func(t *testing.T) {
		a := assertion.DeepEqual[map[string]int]{
			Expected: map[string]int{"a": 1, "b": 2},
			Got:      map[string]int{"a": 1, "b": 2},
		}
		err := a.Check(context.Background())
		expectNoError(t, err)
	})
	t.Run("Unequal maps", func(t *testing.T) {
		a := assertion.DeepEqual[map[string]int]{
			Expected: map[string]int{"a": 1, "b": 2},
			Got:      map[string]int{"a": 1, "b": 3},
		}
		err := a.Check(context.Background())
		expectError(t, err)
	})
	t.Run("Nil slices equal", func(t *testing.T) {
		var s1, s2 []int
		a := assertion.DeepEqual[[]int]{Expected: s1, Got: s2}
		err := a.Check(context.Background())
		expectNoError(t, err)
	})
}

func TestNotEqual(t *testing.T) {
	t.Run("Unequal integers", func(t *testing.T) {
		a := assertion.NotEqual[int]{Unexpected: 42, Got: 99}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "isNotEqual(unexpected: 42, got: 99)")
	})
	t.Run("Equal integers", func(t *testing.T) {
		a := assertion.NotEqual[int]{Unexpected: 42, Got: 42}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), "isNotEqual(unexpected: 42, got: 42)")
	})
	t.Run("Unequal strings", func(t *testing.T) {
		a := assertion.NotEqual[string]{Unexpected: "hello", Got: "world"}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "isNotEqual(unexpected: hello, got: world)")
	})
	t.Run("Equal strings", func(t *testing.T) {
		a := assertion.NotEqual[string]{Unexpected: "hello", Got: "hello"}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), "isNotEqual(unexpected: hello, got: hello)")
	})
	t.Run("Equal zero values", func(t *testing.T) {
		a := assertion.NotEqual[int]{Unexpected: 0, Got: 0}
		err := a.Check(context.Background())
		expectError(t, err)
	})
	t.Run("Unequal booleans", func(t *testing.T) {
		a := assertion.NotEqual[bool]{Unexpected: true, Got: false}
		err := a.Check(context.Background())
		expectNoError(t, err)
	})
	t.Run("Equal booleans", func(t *testing.T) {
		a := assertion.NotEqual[bool]{Unexpected: true, Got: true}
		err := a.Check(context.Background())
		expectError(t, err)
	})
}

func TestNotDeepEqual(t *testing.T) {
	t.Run("Unequal structs", func(t *testing.T) {
		a := assertion.NotDeepEqual[testStruct]{
			Unexpected: testStruct{Name: "foo", Value: 42},
			Got:        testStruct{Name: "bar", Value: 99},
		}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "isNotDeepEqual(unexpected: {foo 42}, got: {bar 99})")
	})
	t.Run("Equal structs", func(t *testing.T) {
		a := assertion.NotDeepEqual[testStruct]{
			Unexpected: testStruct{Name: "foo", Value: 42},
			Got:        testStruct{Name: "foo", Value: 42},
		}
		err := a.Check(context.Background())
		expectError(t, err)
	})
	t.Run("Unequal slices", func(t *testing.T) {
		a := assertion.NotDeepEqual[[]int]{
			Unexpected: []int{1, 2, 3},
			Got:        []int{1, 2, 4},
		}
		err := a.Check(context.Background())
		expectNoError(t, err)
	})
	t.Run("Equal slices", func(t *testing.T) {
		a := assertion.NotDeepEqual[[]int]{
			Unexpected: []int{1, 2, 3},
			Got:        []int{1, 2, 3},
		}
		err := a.Check(context.Background())
		expectError(t, err)
	})
	t.Run("Unequal byte slices", func(t *testing.T) {
		a := assertion.NotDeepEqual[[]byte]{
			Unexpected: []byte{0x01, 0x02, 0x03},
			Got:        []byte{0x01, 0x02, 0x04},
		}
		err := a.Check(context.Background())
		expectNoError(t, err)
	})
	t.Run("Equal byte slices", func(t *testing.T) {
		a := assertion.NotDeepEqual[[]byte]{
			Unexpected: []byte{0x01, 0x02, 0x03},
			Got:        []byte{0x01, 0x02, 0x03},
		}
		err := a.Check(context.Background())
		expectError(t, err)
	})
	t.Run("Unequal maps", func(t *testing.T) {
		a := assertion.NotDeepEqual[map[string]int]{
			Unexpected: map[string]int{"a": 1, "b": 2},
			Got:        map[string]int{"a": 1, "b": 3},
		}
		err := a.Check(context.Background())
		expectNoError(t, err)
	})
	t.Run("Equal maps", func(t *testing.T) {
		a := assertion.NotDeepEqual[map[string]int]{
			Unexpected: map[string]int{"a": 1, "b": 2},
			Got:        map[string]int{"a": 1, "b": 2},
		}
		err := a.Check(context.Background())
		expectError(t, err)
	})
	t.Run("Nil slices equal", func(t *testing.T) {
		var s1, s2 []int
		a := assertion.NotDeepEqual[[]int]{Unexpected: s1, Got: s2}
		err := a.Check(context.Background())
		expectError(t, err)
	})
}
