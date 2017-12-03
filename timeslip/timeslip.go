package timeslip

import (
	"errors"
	"github.com/google/uuid"
	"strings"
	"time"
)

type Slip struct {
	Project   string
	Task      string
	Comment   string
	Started   int
	Logged    int
	Completed int
	Modified  int
	Status    string
	UUID      string
}

func New(name string) (*Slip, error) {
	currentTime := int(time.Now().Unix())
	project, task, err := parseProjectName(name)
	if err != nil {
		return nil, err
	}

	slip := &Slip{
		Project:  project,
		Task:     task,
		Comment:  "New Time Slip",
		Started:  currentTime,
		Modified: currentTime,
		Status:   "started",
		UUID:     uuid.New().String(),
	}

	return slip, nil
}

func parseProjectName(name string) (string, string, error) {
	names := strings.Split(name, ".")

	switch len(names) {
	case 1:
		return names[0], "", nil
	case 2:
		return names[0], names[1], nil
	default:
		return "", "", errors.New("Bad Project/Task name format. Expected 'ProjectName.TaskName' format")
	}
}
