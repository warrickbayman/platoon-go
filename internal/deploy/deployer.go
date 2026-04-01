package deploy

import (
	"fmt"
	"os"
	"os/exec"
	"platoon-go/internal/config"
	"platoon-go/internal/output"
	"platoon-go/internal/ssh"
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
			err := runRemoteCommand(target, commands[c].Command, logPath)
			if err != nil {
				color.Red(commands[c].Command)
				fmt.Printf("error running remote command: %v\n", err)
				os.Exit(2)
			}
		default:
			fmt.Println(color.New(color.FgBlue).Sprint("[LOCAL]  ") + commands[c].Name)
			err := runLocalCommand(commands[c].Command, logPath)
			if err != nil {
				color.Red(commands[c].Command)
				fmt.Printf("error running local command: %v\n", err)
				os.Exit(2)
			}
		}
	}

	Cleanup(target)

	return nil
}

func runLocalCommand(command string, logPath string) error {
	data, err := exec.Command("bash", "-c", command).CombinedOutput()

	output.WriteToFile(logPath, string(data))

	if err != nil {
		output.WriteToFile(logPath, err.Error())
		return fmt.Errorf("failed to run local command: %v", err)
	}
	return nil
}

func runRemoteCommand(target config.TargetConfig, command string, logPath string) error {
	data, err := ssh.RunShell(target, command)

	output.WriteToFile(logPath, string(data))

	if err != nil {
		output.WriteToFile(logPath, err.Error())
		return fmt.Errorf("failed to run shell command: %v", err)
	}

	return nil
}
