package assertion

import (
	"context"
	"fmt"
	"strings"
)

// Substring asserts that V contains Sub
type Substring struct {
	V   string
	Sub string
}

var _ Assertion = Substring{}

func (s Substring) String() string {
	return fmt.Sprintf("isSubstring(v: %q, sub: %q)", s.V, s.Sub)
}

func (s Substring) Check(ctx context.Context) error {
	if !strings.Contains(s.V, s.Sub) {
		return fmt.Errorf("string %s does not contain substring %q", s.V, s.Sub)
	}
	return nil
}

// Prefix asserts that V is prefixed by Prefix
type Prefix struct {
	V      string
	Prefix string
}

var _ Assertion = Prefix{}

func (s Prefix) String() string {
	return fmt.Sprintf("hasPrefix(v: %q, prefix: %q)", s.V, s.Prefix)
}

func (s Prefix) Check(ctx context.Context) error {
	if !strings.HasPrefix(s.V, s.Prefix) {
		return fmt.Errorf("string %s does not have prefix %q", s.V, s.Prefix)
	}
	return nil
}

// Suffix asserts that V is suffixed by Suffix
type Suffix struct {
	V      string
	Suffix string
}

var _ Assertion = Suffix{}

func (s Suffix) String() string {
	return fmt.Sprintf("hasSuffix(v: %q, suffix: %q)", s.V, s.Suffix)
}

func (s Suffix) Check(ctx context.Context) error {
	if !strings.HasSuffix(s.V, s.Suffix) {
		return fmt.Errorf("string %s does not have suffix %q", s.V, s.Suffix)
	}
	return nil
}
