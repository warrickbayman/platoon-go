package cmd

import (
	"fmt"
	"platoon-go/internal/release"
	"platoon-go/internal/targets"

	"github.com/spf13/cobra"
)

var releaseActivateCmd = &cobra.Command{
	Use:     "release:activate [id]",
	Short:   "Activate the specified release",
	Long:    "Activate an existing release on the specified target. The release MUST exist on the target server",
	GroupID: "releases",
	Args:    cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := loadConfig()

		target := targets.ResolveTarget(cfg, resolveTargetName(cfg, targetName))

		fmt.Println("Activating release on " + target.Host)

		release.Activate(target, args[0])

		return nil
	},
}

func init() {
	releaseActivateCmd.Flags().StringVarP(&targetName, "target", "t", "", "Target name")
}
