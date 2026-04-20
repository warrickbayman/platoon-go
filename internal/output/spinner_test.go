package output

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/fatih/color"
)

func captureStdout(fn func()) string {
	r, w, _ := os.Pipe()
	old := os.Stdout
	os.Stdout = w

	fn()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

func TestNewSpinner(t *testing.T) {
	s := NewSpinner("loading")
	if s.label != "loading" {
		t.Errorf("label = %q; want %q", s.label, "loading")
	}
	if s.stop == nil {
		t.Error("stop channel is nil")
	}
	if s.done == nil {
		t.Error("done channel is nil")
	}
}

func TestSpinnerSuccess(t *testing.T) {
	color.NoColor = false
	s := NewSpinner("deploying")

	out := captureStdout(func() {
		s.Start()
		s.Success()
	})

	want := color.New(color.FgGreen).Sprint("✓")
	if !strings.Contains(out, want) {
		t.Errorf("output %q does not contain success mark %q", out, want)
	}
	if !strings.Contains(out, "deploying") {
		t.Errorf("output %q does not contain label %q", out, "deploying")
	}
}

func TestSpinnerFail(t *testing.T) {
	color.NoColor = false
	s := NewSpinner("deploying")

	out := captureStdout(func() {
		s.Start()
		s.Fail()
	})

	want := color.New(color.FgRed).Sprint("✗")
	if !strings.Contains(out, want) {
		t.Errorf("output %q does not contain fail mark %q", out, want)
	}
	if !strings.Contains(out, "deploying") {
		t.Errorf("output %q does not contain label %q", out, "deploying")
	}
}
