package release

import (
	"os"
	"platoon-go/internal/config"
	"platoon-go/internal/targets"

	"github.com/spf13/cobra"
)

var configFilename string
var cfg *config.Config
var targetName string
var target *config.TargetConfig

var ReleaseCmd = &cobra.Command{
	Use:     "release",
	Aliases: []string{"r", "rel"},
	Short:   "Manage releases",
	Long:    "List, activate and cleanup existing releases on the target host",
	GroupID: "releases",
	RunE: func(cmd *cobra.Command, args []string) error {

		if len(args) == 0 {
			_ = cmd.Help()
			os.Exit(0)
		}

		return nil
	},
}

func init() {
	ReleaseCmd.PersistentFlags().StringVarP(&configFilename, "config", "c", "platoon.yml", "Path to the platoon config file")
	ReleaseCmd.PersistentFlags().StringVarP(&targetName, "target", "t", "", "The name of the target host")

	ReleaseCmd.AddCommand(listCmd)
	ReleaseCmd.AddCommand(activateCmd)
	ReleaseCmd.AddCommand(cleanupCmd)
}

func loadConfig() {
	cfg = config.Load(configFilename)

	target = targets.ResolveTarget(cfg, targets.ResolveTargetName(cfg, targetName))
}
