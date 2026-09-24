package assertion

import (
	"context"
	"fmt"
)

// Annotated wraps the given assertion with a message and format arguments.
// The formatted message is appended to the error of a failed check.
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
	err := a.Inner.Check(ctx)
	if err != nil && a.Msg != "" {
		return fmt.Errorf("%w: %s", err, fmt.Sprintf(a.Msg, a.Args...))
	}
	return err
}
