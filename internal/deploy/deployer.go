package deploy

import (
	"fmt"
	"os"
	"platoon-go/internal/config"
	"platoon-go/internal/output"
	"platoon-go/internal/shell"
	"time"

	"github.com/fatih/color"
)

func Run(target config.TargetConfig, gitRepo string, logPath string) error {
	fmt.Println()
	color.Green("Deploy to " + target.Host + "...")

	releaseId := time.Now().Format("20060102150405")

	fmt.Println("Release ID: " + color.New(color.FgBlue).Sprint(releaseId))

	commands := BuildCommands(target, gitRepo, releaseId)

	os.Remove(logPath)

	for c := range commands {

		if commands[c].Command == "" {
			continue
		}

		output.WriteToFile(logPath, commands[c].Name)

		switch commands[c].Type {
		case "remote":
			fmt.Println(color.New(color.FgCyan).Sprint("[REMOTE] ") + commands[c].Name)
			_, err := shell.RunRemoteCommand(target, commands[c].Command)

			output.WriteToFile(logPath, commands[c].Command)

			if err != nil {
				output.WriteToFile(logPath, err.Error())

				color.Red(commands[c].Command)
				fmt.Printf("error running remote command: %v\n", err)
				os.Exit(2)
			}
		default:
			fmt.Println(color.New(color.FgBlue).Sprint("[LOCAL]  ") + commands[c].Name)
			_, err := shell.RunLocalCommand(commands[c].Command)

			output.WriteToFile(logPath, commands[c].Command)

			if err != nil {
				output.WriteToFile(logPath, err.Error())

				color.Red(commands[c].Command)
				fmt.Printf("error running local command: %v\n", err)
				os.Exit(2)
			}
		}
	}

	Cleanup(target)

	return nil
}
