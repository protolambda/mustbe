package mustbe_test

import (
	"context"
	"testing"

	"github.com/protolambda/mustbe"
	"github.com/protolambda/mustbe/assertion"
	"github.com/protolambda/mustbe/be"
)

type testT struct {
	errArgs []any
	failed  bool
}

func (t *testT) Error(args ...any) {
	t.errArgs = args
}

func (t *testT) FailNow() {
	t.failed = true
}

func (t *testT) Context() context.Context {
	return context.Background()
}

func (t *testT) Helper() {

}

func (t *testT) Must(c assertion.Assertion) {
	mustbe.Must(t, c)
}

func (t *testT) Mustf(a assertion.Assertion, msg string, args ...any) {
	mustbe.Must(t, assertion.Annotated{
		Inner: a,
		Msg:   msg,
		Args:  args,
	})
}

var _ mustbe.T = (*testT)(nil)

func TestMust__Success(t *testing.T) {
	got := 123
	expected := 123
	tt := &testT{}
	tt.Mustf(be.Equal(expected, got), "%d", 123)
	t.Helper()

	if tt.failed {
		t.Error("expected success")
	}
	if tt.errArgs != nil {
		t.Error("expected no error")
	}
}

func TestMust__Failure(gt *testing.T) {
	got := 123
	expected := 42
	tt := &testT{}
	tt.Mustf(be.Equal(expected, got), "%d", 123)
	gt.Helper()

	if !tt.failed {
		gt.Error("expected failnow")
	}
	if tt.errArgs == nil {
		gt.Error("expected error")
	}
}

type panicAssertion struct{}

var _ assertion.Assertion = panicAssertion{}

func (p panicAssertion) String() string {
	return "panicAssertion"
}

func (p panicAssertion) Check(ctx context.Context) error {
	panic("surprise")
}

func TestMust__Panic(t *testing.T) {
	tt := &testT{}
	tt.Must(panicAssertion{})
	t.Helper()

	if !tt.failed {
		t.Error("expected failnow")
	}
	if tt.errArgs == nil {
		t.Error("expected error")
	}
	if tt.errArgs[0] != "panic in assertion" {
		t.Error("expected panic in assertion, got ", tt.errArgs[0])
	}
	if tt.errArgs[1] != "surprise" {
		t.Error("expected surprise, got ", tt.errArgs[1])
	}
}
