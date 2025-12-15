package be

import "github.com/protolambda/mustbe/assertion"

func Substring(v string, sub string) assertion.Assertion {
	return assertion.Substring{
		V:   v,
		Sub: sub,
	}
}

func Prefix(v string, prefix string) assertion.Assertion {
	return assertion.Prefix{
		V:      v,
		Prefix: prefix,
	}
}

func Suffix(v string, suffix string) assertion.Assertion {
	return assertion.Suffix{
		V:      v,
		Suffix: suffix,
	}
}
