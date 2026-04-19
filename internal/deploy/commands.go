package deploy

import (
	"platoon-go/internal/config"
	"platoon-go/internal/release"
	"strings"
)

type PlatoonCommand struct {
	Case    func() bool
	Type    string
	Name    string
	Command string
}

func BuildCommands(target *config.TargetConfig, gitRepo string, releaseId string) []*PlatoonCommand {

	return []*PlatoonCommand{
		{
			Case: func() bool {
				return true
			},
			Type:    "local",
			Name:    "Local pre-deploy",
			Command: strings.Join(target.Scripts.LocalPreDeploy, " && "),
		},
		{
			Case: func() bool {
				return true
			},
			Type:    "remote",
			Name:    "Remote pre-deploy",
			Command: strings.Join(target.Scripts.RemotePreDeploy, " && "),
		},
		{
			Case: func() bool {
				return true
			},
			Type:    "remote",
			Name:    "Cloning Git Repo",
			Command: "git clone -b " + target.Branch + " " + gitRepo + " " + target.ReleasePath(releaseId),
		},
		{
			Case: func() bool {
				return !release.FileExists(target, target.StoragePath(""))
			},
			Type:    "remote",
			Name:    "Copying storage directory",
			Command: "if [ ! -d \"" + target.Root + "/storage\" ]; then cp -r " + target.ReleasePath(releaseId+"/storage") + " " + target.Root + "/storage; fi",
		},
		{
			Case: func() bool {
				return true
			},
			Type:    "remote",
			Name:    "Sym-linking storage",
			Command: "rm -rf " + target.ReleasePath(releaseId+"/storage") + " && ln -nfs " + target.StoragePath("") + " " + target.ReleasePath(releaseId+"/storage"),
		},
		{
			Case: func() bool {
				return !release.FileExists(target, target.Root+"/.env")
			},
			Type:    "remote",
			Name:    "Copying .env",
			Command: "cp " + target.ReleasePath(releaseId+"/.env.example") + " " + target.Root + "/.env",
		},
		{
			Case: func() bool {
				return true
			},
			Type:    "remote",
			Name:    "Sym-linking .env",
			Command: "ln -nfs " + target.Root + "/.env " + target.ReleasePath(releaseId+"/.env"),
		},
		{
			Case: func() bool {
				return true
			},
			Type:    "local",
			Name:    "Local post-deploy",
			Command: strings.Join(target.Scripts.LocalPostDeploy, " && "),
		},
		{
			Case: func() bool {
				return len(target.Scripts.RemotePostDeploy) > 0
			},
			Type:    "remote",
			Name:    "Remote post-deploy",
			Command: strings.Join(target.Scripts.RemotePostDeploy, " && "),
		},
		{
			Case: func() bool {
				return true
			},
			Type:    "remote",
			Name:    "Going Live",
			Command: "ln -nfs " + target.ReleasePath(releaseId) + " " + target.ServePath(""),
		},
		{
			Case: func() bool {
				return len(target.Scripts.LocalPostLive) > 0
			},
			Type:    "local",
			Name:    "Local post-live",
			Command: strings.Join(target.Scripts.LocalPostLive, " && "),
		},
		{
			Case: func() bool {
				return len(target.Scripts.RemotePostLive) > 0
			},
			Type:    "remote",
			Name:    "Remote post-live",
			Command: strings.Join(target.Scripts.RemotePostLive, " && "),
		},
	}
}
