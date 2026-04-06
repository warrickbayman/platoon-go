package config

import (
	"fmt"
	"os"
	"platoon-go/internal/output"
	"strings"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Repo    string                  `yaml:"repo"`
	Default string                  `yaml:"default"`
	Targets map[string]TargetConfig `yaml:"targets"`
}

type TargetConfig struct {
	Host     string              `yaml:"host,omitempty"`
	Port     int                 `yaml:"port,omitempty"`
	Username string              `yaml:"username,omitempty"`
	Root     string              `yaml:"root,omitempty"`
	Branch   string              `yaml:"branch,omitempty"`
	Paths    TargetPathConfig    `yaml:"paths,omitempty"`
	Assets   []string            `yaml:"assets,omitempty"`
	Scripts  TargetScriptConfig  `yaml:"scripts,omitempty"`
	Releases TargetReleaseConfig `yaml:"releases,omitempty"`
}

type TargetPathConfig struct {
	Releases string `yaml:"releases,omitempty"`
	Serve    string `yaml:"serve,omitempty"`
	Storage  string `yaml:"storage,omitempty"`
}

type TargetScriptConfig struct {
	LocalPreDeploy   []string `yaml:"local_pre_deploy,omitempty"`
	RemotePreDeploy  []string `yaml:"remote_pre_deploy,omitempty"`
	LocalPostDeploy  []string `yaml:"local_post_deploy,omitempty"`
	RemotePostDeploy []string `yaml:"remote_post_deploy,omitempty"`
	LocalPostLive    []string `yaml:"local_post_live,omitempty"`
	RemotePostLive   []string `yaml:"remote_post_live,omitempty"`
}

type TargetReleaseConfig struct {
	Max int `yaml:"max"`
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

func Load(filename string) *Config {
	fmt.Print("Using config file " + output.Highlight(filename) + "\n\n")

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(output.Error("error loading config: " + err.Error()))
		os.Exit(1)
	}

	var cfg Config

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		fmt.Println(output.Error("error loading config: " + err.Error()))
		os.Exit(1)
	}

	return &cfg
}

func Init(filename string, overwrite bool) {
	cfg := Config{
		Repo:    "git@github.com:team/repo.git",
		Default: "staging",
		Targets: map[string]TargetConfig{
			"common": {
				Host:     "common.host",
				Port:     22,
				Username: "user",
			},
			"staging": {
				Root:   "/path/to/project/root",
				Branch: "main",
				Assets: []string{
					"public/build:public/build",
				},
				Scripts: TargetScriptConfig{
					LocalPreDeploy: []string{
						"npm i",
						"npm run build",
					},
				},
			},
		},
	}

	data, _ := yaml.Marshal(cfg)
	_, err := os.Stat(filename)
	if os.IsNotExist(err) || !overwrite {
		fmt.Println(output.Error("config file already exists. Use --force to overwrite."))
		os.Exit(1)
	}

	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(output.Error("error creating config file: " + err.Error()))
		os.Exit(1)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println(output.Error("error closing config file: " + err.Error()))
			os.Exit(1)
		}
	}(f)

	_, er := fmt.Fprint(f, string(data))

	if er != nil {
		fmt.Println(output.Error("error writing to config file: " + er.Error()))
		os.Exit(1)
	}
}
