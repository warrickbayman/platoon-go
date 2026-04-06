package release

import (
	"errors"
	"platoon-go/internal/config"
	"platoon-go/internal/shell"
	"slices"
	"strings"
	"time"
)

type Release struct {
	Id     string
	Date   string
	Active bool
}

func Activate(target *config.TargetConfig, id string) error {

	_, err := shell.RunRemoteCommand(target, "ln -nfs "+target.ReleasePath(id)+" "+target.ServePath(""))

	if err != nil {
		return err
	}

	active, _ := Active(target)

	if active.Id != id {
		return errors.New("unable to set release")
	}

	return nil
}

func List(target *config.TargetConfig) ([]Release, error) {

	data, err := shell.RunRemoteCommand(target, "ls "+target.ReleasePath("/"))

	if err != nil {
		return nil, err
	}

	active, err := Active(target)

	ids := strings.Split(data, "\n")
	ids = slices.Delete(ids, len(ids)-1, len(ids))

	releases := make([]Release, len(ids))
	for i, id := range ids {
		releases[i] = Release{
			Id:     id,
			Date:   DateString(id),
			Active: active.Id == id,
		}
	}

	return releases, nil
}

func Active(target *config.TargetConfig) (*Release, error) {
	data, err := shell.RunRemoteCommand(target, "ls -la "+target.Root+"/live")

	if err != nil {
		return nil, err
	}

	link := strings.Split(data, " -> ")

	if len(link) != 2 {
		return nil, errors.New("no active release found")
	}

	pieces := strings.Split(link[1], "/")
	id := pieces[len(pieces)-1]

	id = strings.Trim(id, "\n")

	return &Release{
		Id:     id,
		Date:   DateString(id),
		Active: true,
	}, nil
}

func DateString(id string) string {
	t, _ := time.Parse("20060102150405", id)
	return t.Format(time.DateTime)
}
