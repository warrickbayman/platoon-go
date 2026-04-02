package release

import (
	"fmt"
	"platoon-go/internal/config"
	"platoon-go/internal/shell"
)

func List(target config.TargetConfig) (string, error) {

	data, err := shell.RunRemoteCommand(target, "ls -l")

	if err != nil {
		return "", err
	}

	fmt.Printf(data)

	return string(data), nil
}
