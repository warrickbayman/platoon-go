package cmd

import (
	"fmt"
	"os"
	"platoon-go/internal/cli/config"
	"platoon-go/internal/cli/deploy"
	"platoon-go/internal/cli/release"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var Version string

var rootCmd = &cobra.Command{
	Use:   "platoon",
	Short: "Zero-download deployments",
	Long:  "Platoon is a CLI deployment tool that helps make zero-downtime deployments simple",
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flag("version").Changed {
			showVersion()
			os.Exit(0)
		}

		err := cmd.Help()

		if err != nil {
			fmt.Println(err)
		}
	},
}

func Execute(version string) {
	Version = version

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}

func init() {
	rootCmd.Flags().BoolP("version", "v", false, "Show version information")

	rootCmd.AddCommand(config.InitCmd)
	rootCmd.AddCommand(deploy.DeployCmd)

	rootCmd.AddGroup(&cobra.Group{ID: "releases", Title: "Release Management"})
	rootCmd.AddCommand(release.ReleaseCmd)
}

func showVersion() {
	fmt.Println("Platoon-Go : " + color.New(color.FgCyan).Sprint(Version))
}
