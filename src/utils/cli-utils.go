package utils

import (
	"errors"
	"strings"
)

type TaskId struct {
	ProjectName string
	TaskName    string
}

const defaultProjectName = "Global"

func ParseJobId(raw string) (*TaskId, error) {
	parsed := strings.Split(raw, ":")
	if len(parsed) < 1 {
		return nil, errors.New("invalid tas")
	} else if len(parsed) < 2 {
		return &TaskId{
			ProjectName: defaultProjectName,
			TaskName:    parsed[0],
		}, nil
	} else {
		return &TaskId{
			ProjectName: strings.Trim(parsed[0], " "),
			TaskName:    strings.Trim(parsed[1], " "),
		}, nil
	}
}
