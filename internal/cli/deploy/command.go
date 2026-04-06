package deploy

import (
	"platoon-go/internal/config"
	"platoon-go/internal/deploy"
	"platoon-go/internal/targets"

	"github.com/spf13/cobra"
)

var configFilename string
var logFilename string
var targetName string

var DeployCmd = &cobra.Command{
	Use:     "deploy",
	Aliases: []string{"d", "dep"},
	Short:   "Run to a target",
	Long:    "Run the application to the specified target (or default target if none is specified)",
	Args:    cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := config.Load(configFilename)

		target := targets.ResolveTarget(cfg, targets.ResolveTargetName(cfg, targetName))

		err := deploy.Run(target, cfg.Repo, logFilename)

		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	DeployCmd.Flags().StringVarP(&configFilename, "config", "c", "platoon.yml", "Path to a platoon config file")
	DeployCmd.Flags().StringVarP(&logFilename, "log", "l", "deploy.log", "Log file to log to")
	DeployCmd.Flags().StringVarP(&targetName, "name", "n", "", "Target name")
}
