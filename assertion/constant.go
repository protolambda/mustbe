package assertion

import (
	"context"
	"errors"
)

// Failed is a pseudo-assertion that always fails, for placeholder usage.
type Failed struct{}

var _ Assertion = Failed{}

func (a Failed) String() string {
	return "failed"
}

func (a Failed) Check(ctx context.Context) error {
	return errors.New("failed")
}

// Passed is a pseudo-assertion that always passes, for placeholder usage.
type Passed struct{}

var _ Assertion = Passed{}

func (a Passed) String() string {
	return "passed"
}

func (a Passed) Check(ctx context.Context) error {
	return nil
}
