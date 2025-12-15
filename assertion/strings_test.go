package assertion_test

import (
	"context"
	"testing"

	"github.com/protolambda/mustbe/assertion"
)

func TestSubstring(t *testing.T) {
	t.Run("Contains substring", func(t *testing.T) {
		a := assertion.Substring{V: "hello world", Sub: "world"}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), `isSubstring(v: "hello world", sub: "world")`)
	})
	t.Run("Does not contain substring", func(t *testing.T) {
		a := assertion.Substring{V: "hello world", Sub: "foo"}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), `isSubstring(v: "hello world", sub: "foo")`)
	})
	t.Run("Substring at beginning", func(t *testing.T) {
		a := assertion.Substring{V: "hello world", Sub: "hello"}
		err := a.Check(context.Background())
		expectNoError(t, err)
	})
	t.Run("Substring at end", func(t *testing.T) {
		a := assertion.Substring{V: "hello world", Sub: "world"}
		err := a.Check(context.Background())
		expectNoError(t, err)
	})
	t.Run("Empty substring", func(t *testing.T) {
		a := assertion.Substring{V: "hello", Sub: ""}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), `isSubstring(v: "hello", sub: "")`)
	})
	t.Run("Empty string with empty substring", func(t *testing.T) {
		a := assertion.Substring{V: "", Sub: ""}
		err := a.Check(context.Background())
		expectNoError(t, err)
	})
	t.Run("Empty string with non-empty substring", func(t *testing.T) {
		a := assertion.Substring{V: "", Sub: "foo"}
		err := a.Check(context.Background())
		expectError(t, err)
	})
	t.Run("Exact match", func(t *testing.T) {
		a := assertion.Substring{V: "hello", Sub: "hello"}
		err := a.Check(context.Background())
		expectNoError(t, err)
	})
	t.Run("Case sensitive", func(t *testing.T) {
		a := assertion.Substring{V: "Hello World", Sub: "hello"}
		err := a.Check(context.Background())
		expectError(t, err)
	})
}

func TestPrefix(t *testing.T) {
	t.Run("Has prefix", func(t *testing.T) {
		a := assertion.Prefix{V: "hello world", Prefix: "hello"}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), `hasPrefix(v: "hello world", prefix: "hello")`)
	})
	t.Run("Does not have prefix", func(t *testing.T) {
		a := assertion.Prefix{V: "hello world", Prefix: "world"}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), `hasPrefix(v: "hello world", prefix: "world")`)
	})
	t.Run("Empty prefix", func(t *testing.T) {
		a := assertion.Prefix{V: "hello", Prefix: ""}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), `hasPrefix(v: "hello", prefix: "")`)
	})
	t.Run("Empty string with empty prefix", func(t *testing.T) {
		a := assertion.Prefix{V: "", Prefix: ""}
		err := a.Check(context.Background())
		expectNoError(t, err)
	})
	t.Run("Empty string with non-empty prefix", func(t *testing.T) {
		a := assertion.Prefix{V: "", Prefix: "foo"}
		err := a.Check(context.Background())
		expectError(t, err)
	})
	t.Run("Exact match", func(t *testing.T) {
		a := assertion.Prefix{V: "hello", Prefix: "hello"}
		err := a.Check(context.Background())
		expectNoError(t, err)
	})
	t.Run("Case sensitive", func(t *testing.T) {
		a := assertion.Prefix{V: "Hello World", Prefix: "hello"}
		err := a.Check(context.Background())
		expectError(t, err)
	})
	t.Run("Prefix longer than string", func(t *testing.T) {
		a := assertion.Prefix{V: "hi", Prefix: "hello"}
		err := a.Check(context.Background())
		expectError(t, err)
	})
}

func TestSuffix(t *testing.T) {
	t.Run("Has suffix", func(t *testing.T) {
		a := assertion.Suffix{V: "hello world", Suffix: "world"}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), `hasSuffix(v: "hello world", suffix: "world")`)
	})
	t.Run("Does not have suffix", func(t *testing.T) {
		a := assertion.Suffix{V: "hello world", Suffix: "hello"}
		err := a.Check(context.Background())
		expectError(t, err)
		expectString(t, a.String(), `hasSuffix(v: "hello world", suffix: "hello")`)
	})
	t.Run("Empty suffix", func(t *testing.T) {
		a := assertion.Suffix{V: "hello", Suffix: ""}
		err := a.Check(context.Background())
		expectNoError(t, err)
		expectString(t, a.String(), `hasSuffix(v: "hello", suffix: "")`)
	})
	t.Run("Empty string with empty suffix", func(t *testing.T) {
		a := assertion.Suffix{V: "", Suffix: ""}
		err := a.Check(context.Background())
		expectNoError(t, err)
	})
	t.Run("Empty string with non-empty suffix", func(t *testing.T) {
		a := assertion.Suffix{V: "", Suffix: "foo"}
		err := a.Check(context.Background())
		expectError(t, err)
	})
	t.Run("Exact match", func(t *testing.T) {
		a := assertion.Suffix{V: "hello", Suffix: "hello"}
		err := a.Check(context.Background())
		expectNoError(t, err)
	})
	t.Run("Case sensitive", func(t *testing.T) {
		a := assertion.Suffix{V: "Hello World", Suffix: "WORLD"}
		err := a.Check(context.Background())
		expectError(t, err)
	})
	t.Run("Suffix longer than string", func(t *testing.T) {
		a := assertion.Suffix{V: "hi", Suffix: "hello"}
		err := a.Check(context.Background())
		expectError(t, err)
	})
}
