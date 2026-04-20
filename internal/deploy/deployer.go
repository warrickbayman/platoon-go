package deploy

import (
	"fmt"
	"os"
	"platoon-go/internal/config"
	"platoon-go/internal/output"
	"platoon-go/internal/release"
	"platoon-go/internal/shell"
	"time"

	"github.com/fatih/color"
)

func Run(target *config.TargetConfig, gitRepo string, logPath string) error {
	fmt.Println()
	color.Green("Deploy to " + target.Host + "...")

	releaseId := time.Now().Format("20060102150405")

	fmt.Println("Release ID: " + output.Note(releaseId))

	commands := BuildCommands(target, gitRepo, releaseId)

	err := os.Remove(logPath)
	if err != nil {
		fmt.Println(output.Danger("error clearing log file"))
		output.WriteToFile(logPath, err.Error())
		os.Exit(1)
	}

	for c := range commands {

		if commands[c].Command == "" || !commands[c].Case() {
			continue
		}

		output.WriteToFile(logPath, commands[c].Name)

		output.WriteToFile(logPath, commands[c].Command)

		switch commands[c].Type {
		case "remote":
			spinner := output.NewSpinner(output.DarkEmphasis("[REMOTE] ") + commands[c].Name)
			spinner.Start()

			commandOutput, err := shell.RunRemoteCommand(target, commands[c].Command)

			output.WriteToFile(logPath, commandOutput)

			if err != nil {
				spinner.Fail()
				output.WriteToFile(logPath, err.Error())

				color.Red(commands[c].Command)
				fmt.Printf("error running remote command: %v\n", err)
				os.Exit(2)
			}
			spinner.Success()
		default:
			spinner := output.NewSpinner(output.Emphasis("[LOCAL]  ") + commands[c].Name)
			spinner.Start()
			commandOutput, err := shell.RunLocalCommand(commands[c].Command)

			output.WriteToFile(logPath, commandOutput)

			if err != nil {
				spinner.Fail()
				output.WriteToFile(logPath, err.Error())

				color.Red(commands[c].Command)
				fmt.Printf("error running local command: %v\n", err)
				os.Exit(2)
			}
			spinner.Success()
		}
	}

	fmt.Println("----------------------------------------")

	er := release.Cleanup(target)

	if er != nil {
		fmt.Println(output.Danger("error cleaning up releases"))
	}

	return nil
}
