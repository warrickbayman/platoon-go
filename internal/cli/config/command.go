package config

import (
	"platoon-go/internal/config"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "init",
	Aliases: []string{"i", "new"},
	Short:   "Initialize a new platoon config",
	Long:    "Places a new platoon.yml config file at the root of the current project.",
	RunE: func(cmd *cobra.Command, args []string) error {
		configFile, _ := cmd.Flags().GetString("config")
		return config.Init(configFile, cmd.Flags().Changed("force"))
	},
}

func init() {
	Cmd.Flags().StringP("config", "c", "platoon.yml", "Path to the config file")
	Cmd.Flags().BoolP("force", "f", false, "Force overwrite of existing config file")
}
