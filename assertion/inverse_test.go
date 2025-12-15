package assertion_test

import (
	"context"
	"testing"

	"github.com/protolambda/mustbe/assertion"
)

func TestNot(t *testing.T) {
	t.Run("Not of Failed passes", func(t *testing.T) {
		a := assertion.Not{Inner: assertion.Failed{}}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "not(failed)")
	})
	t.Run("Not of Passed fails", func(t *testing.T) {
		a := assertion.Not{Inner: assertion.Passed{}}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), "not(passed)")
	})
	t.Run("Double negation of Failed fails", func(t *testing.T) {
		a := assertion.Not{Inner: assertion.Not{Inner: assertion.Failed{}}}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), "not(not(failed))")
	})
	t.Run("Double negation of Passed passes", func(t *testing.T) {
		a := assertion.Not{Inner: assertion.Not{Inner: assertion.Passed{}}}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "not(not(passed))")
	})
}
