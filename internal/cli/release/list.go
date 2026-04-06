package release

import (
	"fmt"
	"os"
	"platoon-go/internal/output"
	"platoon-go/internal/release"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l", "ls"},
	Short:   "List all existing releases",
	Long:    "List all existing currently installed releases",
	Args:    cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		loadConfig()

		fmt.Println("Releases installed on " + output.Emphasis(target.Host))

		list, err := release.List(target)

		if err != nil {
			return err
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.Header("Release ID", "Date", "Active")

		for _, release := range list {
			isActive := ""
			if release.Active {
				isActive = "*"
			}

			_ = table.Append([]string{
				color.New(color.FgGreen).Sprint(release.Id),
				release.Date,
				color.New(color.FgGreen).Sprint(isActive),
			})
		}

		_ = table.Render()

		return nil
	},
}
