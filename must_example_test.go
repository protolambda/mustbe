package mustbe_test

import (
	"context"
	"fmt"

	"github.com/protolambda/mustbe"
	"github.com/protolambda/mustbe/be"
)

type exampleT struct{}

var _ mustbe.T = (*exampleT)(nil)

func (t *exampleT) Error(args ...any) {
	fmt.Println(append([]any{"ERROR"}, args...)...)
}

func (t *exampleT) FailNow() {
	fmt.Println("FAIL-NOW")
}

func (t *exampleT) Context() context.Context {
	return context.Background()
}

func (t *exampleT) Helper() {}

func ExampleMT_Mustf() {
	t := mustbe.WrapT(new(exampleT))
	t.Mustf(be.Equal(1, 2), "no surprise %d please!", 2)
	// Output:
	// ERROR assertion failed: not equal, expected: 1, got: 2: no surprise 2 please!
	// FAIL-NOW
}

func ExampleMT_Must() {
	t := mustbe.WrapT(new(exampleT))
	t.Must(be.True(false))
	// Output:
	// ERROR assertion failed: expected True
	// FAIL-NOW
}
