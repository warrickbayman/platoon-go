package targets

import (
	"os"
	"platoon-go/internal/config"

	"github.com/fatih/color"
)

func ResolveTarget(cfg *config.Config, name string) *config.TargetConfig {

	t, ok := cfg.Targets[name]

	if !ok {
		color.New(color.FgRed).Print("Target ")
		color.New(color.FgRed, color.Bold).Print(name)
		color.New(color.FgRed).Println(" not found")
		os.Exit(1)
	}

	return &t
}
