package be

import "github.com/protolambda/mustbe/assertion"

// Failed is a pseudo-assertion that always fails, for placeholder usage.
func Failed() assertion.Assertion {
	return assertion.Failed{}
}

// Passed is a pseudo-assertion that always passes, for placeholder usage.
func Passed() assertion.Assertion {
	return assertion.Passed{}
}
