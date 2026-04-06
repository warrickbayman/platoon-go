package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/cobra"
)

func TestInitCmd(t *testing.T) {
	t.Run("Initialize new config", func(t *testing.T) {
		tmpDir := t.TempDir()
		configFile := filepath.Join(tmpDir, "platoon.yml")

		cmd := &cobra.Command{
			RunE: InitCmd.RunE,
		}
		cmd.Flags().StringVarP(&configFile, "config", "c", "platoon.yml", "Path to the config file")
		cmd.Flags().BoolP("force", "f", false, "Force overwrite of existing config file")

		cmd.SetArgs([]string{"--config", configFile})

		err := cmd.Execute()
		if err != nil {
			t.Fatalf("cmd.Execute failed: %v", err)
		}

		if _, err := os.Stat(configFile); os.IsNotExist(err) {
			t.Errorf("Config file was not created at %s", configFile)
		}
	})

	t.Run("Initialize existing config without force", func(t *testing.T) {
		tmpDir := t.TempDir()
		configFile := filepath.Join(tmpDir, "platoon.yml")

		err := os.WriteFile(configFile, []byte("existing"), 0644)
		if err != nil {
			t.Fatalf("Failed to create existing file: %v", err)
		}

		cmd := &cobra.Command{
			RunE: InitCmd.RunE,
		}
		cmd.Flags().StringVarP(&configFile, "config", "c", "platoon.yml", "Path to the config file")
		cmd.Flags().BoolP("force", "f", false, "Force overwrite of existing config file")

		cmd.SetArgs([]string{"--config", configFile})

		err = cmd.Execute()
		if err == nil {
			t.Errorf("Expected error when file exists and force is false, but got nil")
		}
	})

	t.Run("Initialize existing config with force", func(t *testing.T) {
		tmpDir := t.TempDir()
		configFile := filepath.Join(tmpDir, "platoon.yml")

		// Create file first
		err := os.WriteFile(configFile, []byte("existing"), 0644)
		if err != nil {
			t.Fatalf("Failed to create existing file: %v", err)
		}

		cmd := &cobra.Command{
			RunE: InitCmd.RunE,
		}
		cmd.Flags().StringVarP(&configFile, "config", "c", "platoon.yml", "Path to the config file")
		cmd.Flags().BoolP("force", "f", false, "Force overwrite of existing config file")

		cmd.SetArgs([]string{"--config", configFile, "--force"})

		err = cmd.Execute()
		if err != nil {
			t.Fatalf("cmd.Execute failed: %v", err)
		}

		content, err := os.ReadFile(configFile)
		if err != nil {
			t.Fatalf("Failed to read config file: %v", err)
		}

		if string(content) == "existing" {
			t.Errorf("Config file was not overwritten")
		}
	})
}
