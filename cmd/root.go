package cmd

import (
	"fmt"
	"os"
	"platoon-go/internal/config"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

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

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}

func init() {
	rootCmd.Flags().BoolP("version", "v", false, "Show version information")

	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(deployCmd)

	rootCmd.AddGroup(&cobra.Group{ID: "releases", Title: "Release Management"})
	rootCmd.AddCommand(releasesListCmd)
	rootCmd.AddCommand(releaseActivateCmd)
}

func showVersion() {
	fmt.Println("Platoon-Go : " + color.New(color.FgCyan).Sprint("0.0.0-0.1.1"))
}

func resolveTargetName(cfg *config.Config, targetName string) string {
	if targetName == "" {
		return cfg.Default
	}

	return targetName
}
