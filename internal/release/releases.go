package release

import (
	"errors"
	"fmt"
	"platoon-go/internal/config"
	"platoon-go/internal/output"
	"platoon-go/internal/shell"
	"slices"
	"strconv"
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

	active, er := Active(target)

	if er != nil {
		return nil, er
	}

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

func Cleanup(target *config.TargetConfig) error {

	releases, err := List(target)

	if len(releases) == 0 {
		return nil
	}

	if err != nil {
		return err
	}

	slices.SortFunc(releases, func(i, j Release) int {

		iDate, _ := strconv.ParseInt(i.Id, 10, 64)
		jDate, _ := strconv.ParseInt(j.Id, 10, 64)

		if iDate < jDate {
			return 1
		}

		if iDate > jDate {
			return -1
		}

		return 0
	})

	activeIndex := slices.IndexFunc(releases, func(r Release) bool {
		return r.Active
	})

	releasesToClear := make([]Release, 0)

	count := activeIndex + target.Releases.Max

	if len(releases) > count {
		releasesToClear = slices.Delete(releases, 0, activeIndex+target.Releases.Max)
	}

	for _, release := range releasesToClear {
		fmt.Println("Clearing release " + output.Emphasis(release.Id))
		_, err := shell.RunRemoteCommand(target, "rm -rf "+target.ReleasePath(release.Id))

		if err != nil {
			fmt.Println(output.Error("failed to clean release " + release.Id))
		}
	}

	// Clean up old releases
	return nil
}

func DateString(id string) string {
	t, _ := time.Parse("20060102150405", id)
	return t.Format(time.DateTime)
}
