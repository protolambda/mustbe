package assertion_test

import (
	"context"
	"testing"

	"github.com/protolambda/mustbe/assertion"
)

func TestFalse(t *testing.T) {
	t.Run("True value", func(t *testing.T) {
		a := assertion.False{V: true}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), "isFalse(true)")
	})
	t.Run("False value", func(t *testing.T) {
		a := assertion.False{V: false}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "isFalse(false)")
	})
}

func TestTrue(t *testing.T) {
	t.Run("True value", func(t *testing.T) {
		a := assertion.True{V: true}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), "isTrue(true)")
	})
	t.Run("False value", func(t *testing.T) {
		a := assertion.True{V: false}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), "isTrue(false)")
	})
}
