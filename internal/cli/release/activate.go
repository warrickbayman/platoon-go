package release

import (
	"fmt"
	"platoon-go/internal/release"

	"github.com/spf13/cobra"
)

var activateCmd = &cobra.Command{
	Use:     "activate [id]",
	Aliases: []string{"a", "active"},
	Short:   "Activate the specified release",
	Long:    "Activate an existing release on the specified target. The release MUST exist on the target server",
	Args:    cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		loadConfig()

		fmt.Println("Activating release on " + target.Host)

		err := release.Activate(target, args[0])
		if err != nil {
			return fmt.Errorf("unable to activate the release: %w", err)
		}

		return nil
	},
}
