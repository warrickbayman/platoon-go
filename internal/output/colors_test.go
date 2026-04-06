package output

import (
	"github.com/fatih/color"
	"testing"
)

func TestHighlight(t *testing.T) {
	color.NoColor = false // Force color output for testing
	text := "hello"
	expected := color.New(color.FgGreen).Sprint(text)
	result := Highlight(text)
	if result != expected {
		t.Errorf("Highlight(%q) = %q; want %q", text, result, expected)
	}
}

func TestDanger(t *testing.T) {
	color.NoColor = false
	text := "danger"
	expected := color.New(color.FgRed).Sprint(text)
	result := Danger(text)
	if result != expected {
		t.Errorf("Danger(%q) = %q; want %q", text, result, expected)
	}
}

func TestError(t *testing.T) {
	color.NoColor = false
	text := "error"
	expected := color.New(color.FgWhite, color.BgRed).Sprint(text)
	result := Error(text)
	if result != expected {
		t.Errorf("Error(%q) = %q; want %q", text, result, expected)
	}
}

func TestEmphasis(t *testing.T) {
	color.NoColor = false
	text := "emphasis"
	expected := color.New(color.FgCyan).Sprint(text)
	result := Emphasis(text)
	if result != expected {
		t.Errorf("Emphasis(%q) = %q; want %q", text, result, expected)
	}
}

func TestDarkEmphasis(t *testing.T) {
	color.NoColor = false
	text := "dark emphasis"
	expected := color.New(color.FgBlue).Sprint(text)
	result := DarkEmphasis(text)
	if result != expected {
		t.Errorf("DarkEmphasis(%q) = %q; want %q", text, result, expected)
	}
}

func TestNote(t *testing.T) {
	color.NoColor = false
	text := "note"
	expected := color.New(color.FgMagenta).Sprint(text)
	result := Note(text)
	if result != expected {
		t.Errorf("Note(%q) = %q; want %q", text, result, expected)
	}
}
