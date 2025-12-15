package be

import "github.com/protolambda/mustbe/assertion"

// Annotate wraps the given assertion with a message and format arguments.
func Annotate(inner assertion.Assertion, msg string, args ...any) assertion.Annotated {
	return assertion.Annotated{
		Inner: inner,
		Msg:   msg,
		Args:  args,
	}
}
