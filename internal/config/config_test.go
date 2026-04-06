package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestReleasePath(t *testing.T) {
	tests := []struct {
		name     string
		config   TargetConfig
		suffix   string
		expected string
	}{
		{
			name:     "default path with suffix",
			config:   TargetConfig{Root: "/srv/app"},
			suffix:   "20240101",
			expected: "/srv/app/releases/20240101",
		},
		{
			name:     "default path without suffix",
			config:   TargetConfig{Root: "/srv/app"},
			suffix:   "",
			expected: "/srv/app/releases",
		},
		{
			name:     "custom path with suffix",
			config:   TargetConfig{Root: "/srv/app", Paths: TargetPathConfig{Releases: "custom/releases"}},
			suffix:   "20240101",
			expected: "/srv/app/custom/releases/20240101",
		},
		{
			name:     "custom path with trailing slash",
			config:   TargetConfig{Root: "/srv/app", Paths: TargetPathConfig{Releases: "custom/releases/"}},
			suffix:   "20240101",
			expected: "/srv/app/custom/releases/20240101",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.config.ReleasePath(tt.suffix)
			if result != tt.expected {
				t.Errorf("ReleasePath(%q) = %q; want %q", tt.suffix, result, tt.expected)
			}
		})
	}
}

func TestServePath(t *testing.T) {
	tests := []struct {
		name     string
		config   TargetConfig
		suffix   string
		expected string
	}{
		{
			name:     "default path with suffix",
			config:   TargetConfig{Root: "/srv/app"},
			suffix:   "current",
			expected: "/srv/app/live/current",
		},
		{
			name:     "default path without suffix",
			config:   TargetConfig{Root: "/srv/app"},
			suffix:   "",
			expected: "/srv/app/live",
		},
		{
			name:     "custom path with suffix",
			config:   TargetConfig{Root: "/srv/app", Paths: TargetPathConfig{Serve: "public"}},
			suffix:   "current",
			expected: "/srv/app/public/current",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.config.ServePath(tt.suffix)
			if result != tt.expected {
				t.Errorf("ServePath(%q) = %q; want %q", tt.suffix, result, tt.expected)
			}
		})
	}
}

func TestStoragePath(t *testing.T) {
	tests := []struct {
		name     string
		config   TargetConfig
		suffix   string
		expected string
	}{
		{
			name:     "default path with suffix",
			config:   TargetConfig{Root: "/srv/app"},
			suffix:   "uploads",
			expected: "/srv/app/storage/uploads",
		},
		{
			name:     "default path without suffix",
			config:   TargetConfig{Root: "/srv/app"},
			suffix:   "",
			expected: "/srv/app/storage",
		},
		{
			name:     "custom path with suffix",
			config:   TargetConfig{Root: "/srv/app", Paths: TargetPathConfig{Storage: "var/storage"}},
			suffix:   "uploads",
			expected: "/srv/app/var/storage/uploads",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.config.StoragePath(tt.suffix)
			if result != tt.expected {
				t.Errorf("StoragePath(%q) = %q; want %q", tt.suffix, result, tt.expected)
			}
		})
	}
}

func TestLoad(t *testing.T) {
	t.Run("loads valid config", func(t *testing.T) {
		content := `
repo: git@github.com:team/repo.git
default: staging
targets:
  staging:
    host: staging.example.com
    port: 22
    username: deploy
    root: /srv/app
    branch: main
    releases:
      max: 5
`
		f, err := os.CreateTemp(t.TempDir(), "platoon-*.yml")
		if err != nil {
			t.Fatal(err)
		}
		if _, err := f.WriteString(content); err != nil {
			t.Fatal(err)
		}
		f.Close()

		cfg := Load(f.Name())

		if cfg.Repo != "git@github.com:team/repo.git" {
			t.Errorf("Repo = %q; want %q", cfg.Repo, "git@github.com:team/repo.git")
		}
		if cfg.Default != "staging" {
			t.Errorf("Default = %q; want %q", cfg.Default, "staging")
		}
		target, ok := cfg.Targets["staging"]
		if !ok {
			t.Fatal("expected staging target to exist")
		}
		if target.Host != "staging.example.com" {
			t.Errorf("Host = %q; want %q", target.Host, "staging.example.com")
		}
		if target.Port != 22 {
			t.Errorf("Port = %d; want 22", target.Port)
		}
		if target.Branch != "main" {
			t.Errorf("Branch = %q; want %q", target.Branch, "main")
		}
	})
}

func TestInit(t *testing.T) {
	t.Run("creates config file", func(t *testing.T) {
		dir := t.TempDir()
		filename := filepath.Join(dir, "platoon.yml")

		if err := Init(filename, false); err != nil {
			t.Fatalf("Init() error = %v", err)
		}

		if _, err := os.Stat(filename); os.IsNotExist(err) {
			t.Error("expected config file to be created")
		}
	})

	t.Run("returns error if file exists without overwrite", func(t *testing.T) {
		dir := t.TempDir()
		filename := filepath.Join(dir, "platoon.yml")

		if err := Init(filename, false); err != nil {
			t.Fatalf("first Init() error = %v", err)
		}

		if err := Init(filename, false); err == nil {
			t.Error("expected error when file exists and overwrite is false")
		}
	})

	t.Run("overwrites file when overwrite is true", func(t *testing.T) {
		dir := t.TempDir()
		filename := filepath.Join(dir, "platoon.yml")

		if err := Init(filename, false); err != nil {
			t.Fatalf("first Init() error = %v", err)
		}

		if err := Init(filename, true); err != nil {
			t.Errorf("Init() with overwrite=true error = %v", err)
		}
	})

	t.Run("written file parses as valid config", func(t *testing.T) {
		dir := t.TempDir()
		filename := filepath.Join(dir, "platoon.yml")

		if err := Init(filename, false); err != nil {
			t.Fatalf("Init() error = %v", err)
		}

		cfg := Load(filename)
		if cfg.Repo == "" {
			t.Error("expected non-empty repo in generated config")
		}
		if cfg.Default == "" {
			t.Error("expected non-empty default in generated config")
		}
		if len(cfg.Targets) == 0 {
			t.Error("expected at least one target in generated config")
		}
	})
}
