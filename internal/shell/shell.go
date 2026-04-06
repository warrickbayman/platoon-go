package shell

import (
	"fmt"
	"os/exec"
	"platoon-go/internal/config"
	"platoon-go/internal/ssh"
)

func RunLocalCommand(command string) (string, error) {
	data, err := exec.Command("bash", "-c", command).CombinedOutput()

	if err != nil {
		return "", fmt.Errorf("failed to run local command: %v", err)
	}
	return string(data), nil
}

func RunRemoteCommand(target *config.TargetConfig, command string) (string, error) {
	data, err := ssh.RunShell(target, command)

	if err != nil {
		return "", fmt.Errorf("failed to run shell command: %v", err)
	}

	return string(data), nil
}
