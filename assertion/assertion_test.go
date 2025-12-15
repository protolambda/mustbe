package assertion_test

import "testing"

// Just some super minimal test assertion utils,
// to make the tests of the test assertion utils library a bit nicer...

func expectError(t *testing.T, err error) {
	if err == nil {
		t.Helper()
		t.Error("expected error, got nil")
	}
}

func expectNoError(t *testing.T, err error) {
	if err != nil {
		t.Helper()
		t.Error("expected no error, got ", err)
	}
}

func expectString(t *testing.T, got, expected string) {
	if got != expected {
		t.Helper()
		t.Error("expected string ", expected, " but got ", got)
	}
}
