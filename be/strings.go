package be

import "github.com/protolambda/mustbe/assertion"

// Substring asserts that v contains sub.
func Substring(v string, sub string) assertion.Assertion {
	return assertion.Substring{
		V:   v,
		Sub: sub,
	}
}

// Prefix asserts that v is prefixed by prefix.
func Prefix(v string, prefix string) assertion.Assertion {
	return assertion.Prefix{
		V:      v,
		Prefix: prefix,
	}
}

// Suffix asserts that v is suffixed by suffix.
func Suffix(v string, suffix string) assertion.Assertion {
	return assertion.Suffix{
		V:      v,
		Suffix: suffix,
	}
}
