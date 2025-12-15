package assertion_test

import (
	"context"
	"testing"

	"github.com/protolambda/mustbe/assertion"
)

func TestFailed(t *testing.T) {
	a := assertion.Failed{}
	err := a.Check(context.Background())
	expectError(t, err)
	expectString(t, a.String(), "failed")
}

func TestPassed(t *testing.T) {
	a := assertion.Passed{}
	err := a.Check(context.Background())
	expectNoError(t, err)
	expectString(t, a.String(), "passed")
}
