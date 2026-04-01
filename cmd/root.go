package cmd

import (
	"fmt"
	"os"

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

		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	os.Exit(0)
}

func init() {
	rootCmd.Flags().BoolP("version", "v", false, "Show version information")

	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(deployCmd)
}

func showVersion() {
	fmt.Println("Platoon-Go : version 0.0.0-0.0.1")
}
