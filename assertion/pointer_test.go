package assertion_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/protolambda/mustbe/assertion"
)

func TestNil(t *testing.T) {
	t.Run("Nil value", func(t *testing.T) {
		a := assertion.Nil{V: nil}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "isNil(got: <nil>)")
	})
	t.Run("Nil pointer in interface", func(t *testing.T) {
		var p *int
		a := assertion.Nil{V: p}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "isNil(got: <nil>)")
	})
	t.Run("Nil slice", func(t *testing.T) {
		var s []int
		a := assertion.Nil{V: s}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "isNil(got: [])")
	})
	t.Run("Nil map", func(t *testing.T) {
		var m map[string]int
		a := assertion.Nil{V: m}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "isNil(got: map[])")
	})
	t.Run("Nil chan", func(t *testing.T) {
		var ch chan int
		a := assertion.Nil{V: ch}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "isNil(got: <nil>)")
	})
	t.Run("Nil func", func(t *testing.T) {
		var fn func()
		a := assertion.Nil{V: fn}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "isNil(got: <nil>)")
	})
	t.Run("Nil interface", func(t *testing.T) {
		var iface error
		a := assertion.Nil{V: iface}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "isNil(got: <nil>)")
	})
	t.Run("Non-nil value", func(t *testing.T) {
		a := assertion.Nil{V: 42}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), "isNil(got: 42)")
	})
	t.Run("Non-nil pointer", func(t *testing.T) {
		x := 42
		a := assertion.Nil{V: &x}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), fmt.Sprintf("isNil(got: %v)", &x))
	})
	t.Run("Non-nil chan", func(t *testing.T) {
		ch := make(chan int)
		a := assertion.Nil{V: ch}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), fmt.Sprintf("isNil(got: %v)", ch))
	})
	t.Run("Non-nil func", func(t *testing.T) {
		fn := func() {}
		a := assertion.Nil{V: fn}
		err := a.Check(context.Background())
		expectError(t, err)
	})
}

func TestNonNil(t *testing.T) {
	t.Run("Non-nil value", func(t *testing.T) {
		a := assertion.NonNil{V: 42}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "nonNil(got: 42)")
	})
	t.Run("Non-nil pointer", func(t *testing.T) {
		x := 42
		a := assertion.NonNil{V: &x}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), fmt.Sprintf("nonNil(got: %v)", &x))
	})
	t.Run("Non-nil chan", func(t *testing.T) {
		ch := make(chan int)
		a := assertion.NonNil{V: ch}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), fmt.Sprintf("nonNil(got: %v)", ch))
	})
	t.Run("Non-nil func", func(t *testing.T) {
		fn := func() {}
		a := assertion.NonNil{V: fn}
		err := a.Check(context.Background())
		expectNoError(t, err)
	})
	t.Run("Nil value", func(t *testing.T) {
		a := assertion.NonNil{V: nil}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), "nonNil(got: <nil>)")
	})
	t.Run("Nil pointer in interface", func(t *testing.T) {
		var p *int
		a := assertion.NonNil{V: p}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), "nonNil(got: <nil>)")
	})
	t.Run("Nil chan", func(t *testing.T) {
		var ch chan int
		a := assertion.NonNil{V: ch}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), "nonNil(got: <nil>)")
	})
	t.Run("Nil func", func(t *testing.T) {
		var fn func()
		a := assertion.NonNil{V: fn}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), "nonNil(got: <nil>)")
	})
	t.Run("Nil interface", func(t *testing.T) {
		var iface error
		a := assertion.NonNil{V: iface}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), "nonNil(got: <nil>)")
	})
}

func TestNilPtr(t *testing.T) {
	t.Run("Nil pointer", func(t *testing.T) {
		a := assertion.NilPtr[int]{P: nil}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "nilPtr(got: <nil>)")
	})
	t.Run("Non-nil pointer", func(t *testing.T) {
		x := 42
		a := assertion.NilPtr[int]{P: &x}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), fmt.Sprintf("nilPtr(got: %v)", &x))
	})
}

func TestNonNilPtr(t *testing.T) {
	t.Run("Non-nil pointer", func(t *testing.T) {
		x := 42
		a := assertion.NonNilPtr[int]{P: &x}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), fmt.Sprintf("nonNilPtr(got: %v)", &x))
	})
	t.Run("Nil pointer", func(t *testing.T) {
		a := assertion.NonNilPtr[int]{P: nil}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), "nonNilPtr(got: <nil>)")
	})
}

func TestSame(t *testing.T) {
	t.Run("Same pointer", func(t *testing.T) {
		x := 42
		a := assertion.Same[int]{Expected: &x, Got: &x}
		err := a.Check(context.Background())
		expectNoError(t, err)
	})
	t.Run("Different pointers same value", func(t *testing.T) {
		x, y := 42, 42
		a := assertion.Same[int]{Expected: &x, Got: &y}
		err := a.Check(context.Background())
		expectError(t, err)
	})
	t.Run("Different pointers different values", func(t *testing.T) {
		x, y := 42, 99
		a := assertion.Same[int]{Expected: &x, Got: &y}
		err := a.Check(context.Background())
		expectError(t, err)
	})
	t.Run("Both nil pointers", func(t *testing.T) {
		a := assertion.Same[int]{Expected: nil, Got: nil}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "isSame(expected: <nil>, got: <nil>)")
	})
	t.Run("One nil pointer", func(t *testing.T) {
		x := 42
		a := assertion.Same[int]{Expected: &x, Got: nil}
		err := a.Check(context.Background())
		expectError(t, err)
	})
	t.Run("Same struct pointer", func(t *testing.T) {
		s := testStruct{Name: "foo", Value: 42}
		a := assertion.Same[testStruct]{Expected: &s, Got: &s}
		err := a.Check(context.Background())
		expectNoError(t, err)
	})
	t.Run("Different struct pointers same value", func(t *testing.T) {
		s1 := testStruct{Name: "foo", Value: 42}
		s2 := testStruct{Name: "foo", Value: 42}
		a := assertion.Same[testStruct]{Expected: &s1, Got: &s2}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), fmt.Sprintf("isSame(expected: %v, got: %v)", &s1, &s2))
	})
}
