package config

import (
	"platoon-go/internal/config"

	"github.com/spf13/cobra"
)

var configFile string

var InitCmd = &cobra.Command{
	Use:     "init",
	Aliases: []string{"i", "new"},
	Short:   "Initialize a new platoon config",
	Long:    "Places a new platoon.yml config file at the root of the current project.",
	RunE: func(cmd *cobra.Command, args []string) error {

		config.Init(configFile, cmd.Flags().Changed("force"))

		return nil
	},
}

func init() {
	InitCmd.Flags().StringVarP(&configFile, "config", "c", "platoon.yml", "Path to the config file")
	InitCmd.Flags().BoolP("force", "f", false, "Force overwrite of existing config file")
}
