package release

import (
	"bytes"
	"os"
	"path/filepath"
	"platoon-go/internal/config"
	"platoon-go/internal/ssh"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestListCmd(t *testing.T) {
	tmpDir := t.TempDir()
	configFile := filepath.Join(tmpDir, "platoon.yml")

	// Setup a minimal config
	err := os.WriteFile(configFile, []byte(`
repo: git@github.com:team/repo.git
default: staging
targets:
  staging:
    host: staging.example.com
    username: deploy
    root: /var/www/app
`), 0644)
	if err != nil {
		t.Fatalf("failed to create config file: %v", err)
	}

	// Mock SSH Runner
	originalRunner := ssh.Runner
	defer func() { ssh.Runner = originalRunner }()

	ssh.Runner = func(target *config.TargetConfig, command string) ([]byte, error) {
		if strings.Contains(command, "ls /var/www/app/releases/") {
			return []byte("20231010101010\n20231011101010\n"), nil
		}
		if strings.Contains(command, "ls -la /var/www/app/live") {
			return []byte("lrwxrwxrwx 1 deploy deploy 34 Oct 11 10:10 /var/www/app/live -> /var/www/app/releases/20231011101010\n"), nil
		}
		return nil, nil
	}

	// Prepare the command
	cmd := &cobra.Command{
		RunE: listCmd.RunE,
	}
	// Need to initialize persistent flags for the parent command if we test listCmd directly,
	// but listCmd uses loadConfig() which relies on global variables in command.go.
	// We should set those global variables.
	configFilename = configFile
	targetName = "staging"

	// Capture output
	// To capture output from fmt.Println and tablewriter, we'd need to redirect os.Stdout.
	// Let's do that for better verification.
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err = cmd.Execute()

	w.Close()
	os.Stdout = oldStdout

	var out bytes.Buffer
	_, _ = out.ReadFrom(r)
	got := out.String()

	if !strings.Contains(got, "20231010101010") {
		t.Errorf("Output does not contain release 20231010101010. Got: %s", got)
	}
	if !strings.Contains(got, "20231011101010") {
		t.Errorf("Output does not contain release 20231011101010. Got: %s", got)
	}
	if !strings.Contains(got, "*") {
		t.Errorf("Output does not indicate active release. Got: %s", got)
	}
}
