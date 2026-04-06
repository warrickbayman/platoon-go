package output

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestWriteToFile(t *testing.T) {
	tmpDir := t.TempDir()
	logFile := filepath.Join(tmpDir, "test.log")
	message := "test message"

	WriteToFile(logFile, message)

	content, err := os.ReadFile(logFile)
	if err != nil {
		t.Fatalf("Failed to read log file: %v", err)
	}

	got := string(content)
	if !strings.Contains(got, message) {
		t.Errorf("Log content %q does not contain %q", got, message)
	}

	// Verify timestamp format (roughly)
	// Format is [2006-01-02 15:04:05]
	if !strings.HasPrefix(got, "[") || !strings.Contains(got, "]") {
		t.Errorf("Log content %q does not have expected timestamp format", got)
	}

	// Verify append behavior
	message2 := "another message"
	WriteToFile(logFile, message2)

	content, err = os.ReadFile(logFile)
	if err != nil {
		t.Fatalf("Failed to read log file after second write: %v", err)
	}

	got = string(content)
	if !strings.Contains(got, message) || !strings.Contains(got, message2) {
		t.Errorf("Log content %q does not contain both messages", got)
	}
}

func TestClearFile(t *testing.T) {
	tmpDir := t.TempDir()
	logFile := filepath.Join(tmpDir, "test.log")

	// Create a file first
	WriteToFile(logFile, "to be cleared")

	if _, err := os.Stat(logFile); os.IsNotExist(err) {
		t.Fatalf("Log file should exist before clearing")
	}

	ClearFile(logFile)

	if _, err := os.Stat(logFile); !os.IsNotExist(err) {
		t.Errorf("Log file should not exist after clearing")
	}
}
