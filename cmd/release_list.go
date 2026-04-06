package cmd

import (
	"fmt"
	"os"
	"platoon-go/internal/release"
	"platoon-go/internal/targets"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var releasesListCmd = &cobra.Command{
	Use:     "release:list",
	Short:   "List all existing releases",
	Long:    "List all existing currently installed releases",
	GroupID: "releases",
	Args:    cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := loadConfig()

		target := targets.ResolveTarget(cfg, resolveTargetName(cfg, targetName))

		fmt.Println("Releases installed on " + target.Host)

		list, err := release.List(target)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.Header("Release ID", "Date", "Active")

		for _, release := range list {
			isActive := ""
			if release.Active {
				isActive = "*"
			}

			table.Append([]string{
				color.New(color.FgGreen).Sprint(release.Id),
				release.Date,
				color.New(color.FgGreen).Sprint(isActive),
			})
		}

		table.Render()

		return nil
	},
}

func init() {
	releasesListCmd.Flags().StringVarP(&targetName, "target", "t", "", "The name of the target host")
}
