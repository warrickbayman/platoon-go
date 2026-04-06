package deploy

import (
	"platoon-go/internal/config"
	"strings"
)

type PlatoonCommand struct {
	Type    string
	Name    string
	Command string
}

func BuildCommands(target *config.TargetConfig, gitRepo string, releaseId string) []*PlatoonCommand {

	return []*PlatoonCommand{
		{
			Type:    "local",
			Name:    "Local pre-deploy",
			Command: strings.Join(target.Scripts.LocalPreDeploy, " && "),
		},
		{
			Type:    "remote",
			Name:    "Remote pre-deploy",
			Command: strings.Join(target.Scripts.RemotePreDeploy, " && "),
		},
		{
			Type:    "remote",
			Name:    "Cloning Git Repo",
			Command: "git clone -b " + target.Branch + " " + gitRepo + " " + target.ReleasePath(releaseId),
		},
		{
			Type:    "remote",
			Name:    "Copying storage directory",
			Command: "if [ -d \"" + target.Root + "/storage\" ]; then cp -r " + target.StoragePath(releaseId+"/storage") + " " + target.Root + "/storage; fi",
		},
		{
			Type:    "remote",
			Name:    "Sym-linking storage",
			Command: "rm -rf " + target.ReleasePath(releaseId+"/storage") + " && ln -nfs " + target.StoragePath(releaseId) + " " + target.ReleasePath(releaseId+"/storage"),
		},
		{
			Type:    "remote",
			Name:    "Sym-linking .env",
			Command: "ln -nfs " + target.Root + "/.env " + target.ReleasePath(releaseId+"/.env"),
		},
		{
			Type:    "local",
			Name:    "Local post-deploy",
			Command: strings.Join(target.Scripts.LocalPostDeploy, " && "),
		},
		{
			Type:    "remote",
			Name:    "Remote post-deploy",
			Command: strings.Join(target.Scripts.RemotePostDeploy, " && "),
		},
		{
			Type:    "remote",
			Name:    "Going Live",
			Command: "ln -nfs " + target.ReleasePath(releaseId) + " " + target.ServePath(""),
		},
		{
			Type:    "local",
			Name:    "Local post-live",
			Command: strings.Join(target.Scripts.LocalPostLive, " && "),
		},
		{
			Type:    "remote",
			Name:    "Remote post-live",
			Command: strings.Join(target.Scripts.RemotePostLive, " && "),
		},
	}
}
