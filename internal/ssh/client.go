package ssh

import (
	"os/exec"
	"platoon-go/internal/config"
	"strconv"
)

func RunShell(target *config.TargetConfig, command string) ([]byte, error) {
	return exec.Command("ssh", target.Username+"@"+target.Host, "-p"+strconv.Itoa(target.Port), command).CombinedOutput()
}
