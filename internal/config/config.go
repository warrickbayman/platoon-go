package config

import (
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Repo    string                  `yaml:"repo"`
	Default string                  `yaml:"default"`
	Targets map[string]TargetConfig `yaml:"targets"`
}

type TargetPathConfig struct {
	Releases string `yaml:"releases"`
	Serve    string `yaml:"serve"`
	Storage  string `yaml:"storage"`
}

type TargetScriptConfig struct {
	LocalPreDeploy   []string `yaml:"local_pre_deploy"`
	RemotePreDeploy  []string `yaml:"remote_pre_deploy"`
	LocalPostDeploy  []string `yaml:"local_post_deploy"`
	RemotePostDeploy []string `yaml:"remote_post_deploy"`
	LocalPostLive    []string `yaml:"local_post_live"`
	RemotePostLive   []string `yaml:"remote_post_live"`
}

type TargetConfig struct {
	Host     string             `yaml:"host"`
	Port     int                `yaml:"port"`
	Username string             `yaml:"username"`
	Root     string             `yaml:"root"`
	Branch   string             `yaml:"branch"`
	Paths    TargetPathConfig   `yaml:"paths"`
	Assets   []string           `yaml:"assets"`
	Scripts  TargetScriptConfig `yaml:"scripts"`
}

func (config *TargetConfig) ReleasePath(suffix string) string {
	path := config.Paths.Releases
	if path == "" {
		path = "releases"
	}

	return resolvePath(config, path, suffix)
}

func (config *TargetConfig) ServePath(suffix string) string {
	path := config.Paths.Releases
	if path == "" {
		path = "live"
	}

	return resolvePath(config, path, suffix)
}

func (config *TargetConfig) StoragePath(suffix string) string {
	path := config.Paths.Releases
	if path == "" {
		path = "storage"
	}

	return resolvePath(config, path, suffix)
}

func resolvePath(config *TargetConfig, path string, suffix string) string {
	if !strings.HasSuffix(path, "/") && suffix != "" {
		path += "/"
	}

	return config.Root + "/" + path + suffix
}

func Load(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var cfg Config

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
