package mustbe_test

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"testing"

	"github.com/protolambda/mustbe"
	"github.com/protolambda/mustbe/be"
)

const helperChildEnv = "MUSTBE_HELPER_CHILD"

// wantNextLine prints the file:line of the statement after the call,
// for the parent test to match against the location reported by the testing package.
func wantNextLine() {
	_, _, line, _ := runtime.Caller(1)
	fmt.Printf("WANT must_helper_test.go:%d\n", line+1)
}

// TestHelperChild runs intentionally failing assertions.
// It only runs as a subprocess of TestHelperLocation.
func TestHelperChild(gt *testing.T) {
	if os.Getenv(helperChildEnv) != "1" {
		gt.Skip("only runs as subprocess of TestHelperLocation")
	}
	gt.Run("Must", func(gt *testing.T) {
		wantNextLine()
		mustbe.Must(gt, be.True(false))
	})
	gt.Run("MT.Must", func(gt *testing.T) {
		t := mustbe.WrapT(gt)
		wantNextLine()
		t.Must(be.True(false))
	})
	gt.Run("MT.Mustf", func(gt *testing.T) {
		t := mustbe.WrapT(gt)
		wantNextLine()
		t.Mustf(be.True(false), "msg %d", 1)
	})
	gt.Run("Panic", func(gt *testing.T) {
		t := mustbe.WrapT(gt)
		wantNextLine()
		t.Must(panicAssertion{})
	})
}

var (
	wantRe = regexp.MustCompile(`WANT (\S+:\d+)`)
	gotRe  = regexp.MustCompile(`(?m)^\s+(\S+\.go:\d+): (assertion failed|panic in assertion)`)
)

// TestHelperLocation checks that failures are reported at the line of the Must call,
// not inside mustbe, by inspecting the output of the real testing package.
func TestHelperLocation(t *testing.T) {
	if os.Getenv(helperChildEnv) == "1" {
		t.Skip("parent test")
	}
	for _, name := range []string{"Must", "MT.Must", "MT.Mustf", "Panic"} {
		t.Run(name, func(t *testing.T) {
			cmd := exec.Command(os.Args[0], "-test.run=^TestHelperChild$/^"+regexp.QuoteMeta(name)+"$", "-test.count=1")
			cmd.Env = append(os.Environ(), helperChildEnv+"=1")
			out, err := cmd.CombinedOutput()
			if err == nil {
				t.Fatalf("expected child test to fail, output:\n%s", out)
			}
			want := wantRe.FindSubmatch(out)
			got := gotRe.FindSubmatch(out)
			if want == nil || got == nil {
				t.Fatalf("could not find locations in output:\n%s", out)
			}
			if string(want[1]) != string(got[1]) {
				t.Errorf("expected failure reported at %s, got %s, output:\n%s", want[1], got[1], out)
			}
		})
	}
}
