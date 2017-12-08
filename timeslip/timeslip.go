// Copyright (c) 2017 Michael R. Cook

package timeslip

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/mrcook/time_warrior/timeslip/status"
)

type Slip struct {
	Project     string
	Task        string
	Description string
	Started     int
	Worked      int
	Finished    int
	Modified    int
	Status      string
	UUID        string
}

func New(name string) (*Slip, error) {
	currentTime := int(time.Now().Unix())
	project, task, err := parseProjectName(name)
	if err != nil {
		return nil, err
	}

	slip := &Slip{
		Project:     project,
		Task:        task,
		Description: "New Time Slip",
		Started:     currentTime,
		Modified:    currentTime,
		Status:      status.Started(),
		UUID:        uuid.New().String(),
	}

	return slip, nil
}

func NewFromJSON(blob []byte) (*Slip, error) {
	slip := &Slip{}
	err := json.Unmarshal(blob, slip)
	if err != nil {
		return nil, err
	}
	return slip, nil
}

func (s *Slip) Pause() error {
	if s.Status == status.Paused() {
		return fmt.Errorf("Slip is already paused")
	}

	s.Status = status.Paused()
	s.Worked += timeNow() - s.Modified
	s.Modified = timeNow()

	return nil
}

func (s Slip) isPaused() bool {
	if s.Status == status.Paused() {
		return false
	}
	return true
}

func (s *Slip) Resume() {
	s.Status = status.Started()
	s.Modified = timeNow()
}

func (s *Slip) Done(description string) {
	currentTime := timeNow()

	if s.Status == status.Started() {
		s.Worked += currentTime - s.Modified
		s.Finished = currentTime
		s.Modified = s.Finished
	} else {
		s.Finished = s.Modified
		s.Modified = timeNow()
	}

	s.Description = description
	s.Status = status.Completed()
}

func (s Slip) FullName() string {
	if s.Task == "" {
		return s.Project
	}
	return s.Project + "." + s.Task
}

func (s Slip) String() string {
	started := time.Unix(int64(s.Started), 0).Format("2006-01-02 15:04")
	return fmt.Sprintf("%s | Started: %s | Worked: %d mins | Status: %s", s.FullName(), started, s.minutesWorked(), s.Status)
}

func (s Slip) minutesWorked() int {
	return s.Worked / 60
}

func parseProjectName(name string) (string, string, error) {
	names := strings.Split(name, ".")

	switch len(names) {
	case 1:
		return names[0], "", nil
	case 2:
		return names[0], names[1], nil
	default:
		return "", "", fmt.Errorf("Bad Project/Task name format. Expected 'ProjectName.TaskName' format")
	}
}

func timeNow() int {
	return int(time.Now().Unix())
}
