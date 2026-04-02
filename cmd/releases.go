package cmd

import (
	"fmt"
	"platoon-go/internal/release"

	"github.com/spf13/cobra"
)

var releasesListCmd = &cobra.Command{
	Use:   "releases:list [target]",
	Short: "List all existing releases",
	Long:  "List all existing currently installed releases",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := loadConfig()
		target := resolveTarget(cfg, args[0])

		fmt.Println("Releases installed on " + target.Host)

		release.List(target)

		return nil
	},
	GroupID: "releases",
}
