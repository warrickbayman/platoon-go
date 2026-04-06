package release

import (
	"fmt"
	"os"
	"platoon-go/internal/output"
	"platoon-go/internal/release"

	"github.com/spf13/cobra"
)

var cleanupCmd = &cobra.Command{
	Use:     "cleanup",
	Aliases: []string{"c", "clean", "clear"},
	Short:   "Cleanup old releases",
	Long:    "Cleanup the number of specified releases. Will not remove the active release or any releases with dates AFTER the active release",
	RunE: func(cmd *cobra.Command, args []string) error {
		loadConfig()

		fmt.Println("Cleaning up releases on " + output.Emphasis(target.Host))

		err := release.Cleanup(target)
		if err != nil {
			fmt.Println(output.Error("unable to cleanup releases"))
			os.Exit(1)
		}

		return nil
	},
}
