package cmd

import (
	"fmt"
	"os"
	"platoon-go/internal/config"
	"platoon-go/internal/deploy"
	"platoon-go/internal/targets"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var configFilename string
var logFilename string
var targetName string

var deployCmd = &cobra.Command{
	Use:     "deploy",
	Short:   "Run to a target",
	Long:    "Run the application to the specified target (or default target if none is specified)",
	Aliases: []string{"d"},
	Args:    cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := loadConfig()

		target := targets.ResolveTarget(cfg, resolveTargetName(cfg, targetName))

		err := deploy.Run(target, cfg.Repo, logFilename)

		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	deployCmd.Flags().StringVarP(&configFilename, "config", "c", "platoon.yml", "Path to a platoon config file")
	deployCmd.Flags().StringVarP(&logFilename, "log", "l", "deploy.log", "Log file to log to")
	deployCmd.Flags().StringVarP(&targetName, "name", "n", "", "Target name")
}

func loadConfig() *config.Config {

	fmt.Print("Using config file ")
	color.New(color.FgGreen).Printf("%s...\n", configFilename)

	cfg, err := config.Load(configFilename)

	if err != nil {
		color.New(color.FgRed).Println(fmt.Errorf("error loading config: %s", err).Error())
		os.Exit(1)
	}

	return cfg
}
