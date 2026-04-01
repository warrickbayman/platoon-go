package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ReleasesListCmd = &cobra.Command{
	Use:   "releases:list [target]",
	Short: "List all existing releases",
	Long:  "List all existing currently installed releases",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := loadConfig()
		target := resolveTarget(cfg, args[0])

		fmt.Println("Releases installed on " + target.Host)

		return nil
	},
}
