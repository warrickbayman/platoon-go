package ssh

import (
	"os/exec"
	"platoon-go/internal/config"
	"strconv"
)

var Runner = func(target *config.TargetConfig, command string) ([]byte, error) {
	return exec.Command("ssh", target.Username+"@"+target.Host, "-p"+strconv.Itoa(target.Port), command).CombinedOutput()
}

func RunShell(target *config.TargetConfig, command string) ([]byte, error) {
	return Runner(target, command)
}
