package targets

import (
	"fmt"
	"os"
	"platoon-go/internal/config"
	"platoon-go/internal/output"
)

func ResolveTarget(cfg *config.Config, name string) *config.TargetConfig {

	t, ok := cfg.Targets[name]

	if !ok {
		fmt.Println(output.Error("target " + name + " not found"))
		os.Exit(1)
	}

	common, _ := cfg.Targets["common"]

	return applyCommon(&t, &common)
}

func applyCommon(target *config.TargetConfig, common *config.TargetConfig) *config.TargetConfig {

	if target.Host == "" {
		target.Host = common.Host
	}

	if target.Port == 0 {
		target.Port = common.Port
	}

	if target.Username == "" {
		target.Username = common.Username
	}

	if target.Root == "" {
		target.Root = common.Root
	}

	if target.Branch == "" {
		target.Branch = common.Branch
	}

	if len(target.Assets) == 0 {
		target.Assets = common.Assets
	}

	applyCommonScriptConfig(target, common)
	applyCommonReleaseConfig(target, common)

	return target
}

func applyCommonScriptConfig(target *config.TargetConfig, common *config.TargetConfig) {
	if len(target.Scripts.LocalPreDeploy) == 0 {
		target.Scripts.LocalPreDeploy = common.Scripts.LocalPreDeploy
	}
	if len(target.Scripts.LocalPostDeploy) == 0 {
		target.Scripts.LocalPostDeploy = common.Scripts.LocalPostDeploy
	}
}

func applyCommonReleaseConfig(target *config.TargetConfig, common *config.TargetConfig) {
	if target.Releases.Max == 0 {
		target.Releases.Max = common.Releases.Max
	}
}

func ResolveTargetName(cfg *config.Config, targetName string) string {
	if targetName == "" {
		return cfg.Default
	}

	return targetName
}
