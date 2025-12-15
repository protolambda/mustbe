package assertion

import (
	"context"
	"fmt"
)

// Annotated wraps the given assertion with a message and format arguments.
type Annotated struct {
	Inner Assertion
	Msg   string
	Args  []any
}

var _ Assertion = Annotated{}

func (a Annotated) String() string {
	return fmt.Sprintf(a.Msg, a.Args...) + ": " + a.Inner.String()
}

func (a Annotated) Check(ctx context.Context) error {
	return a.Inner.Check(ctx)
}
