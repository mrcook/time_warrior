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
	Project     string `json:"project"`
	Task        string `json:"task"`
	Description string `json:"description"`
	Started     int    `json:"started"`
	Worked      int    `json:"worked"`
	Finished    int    `json:"finished"`
	Modified    int    `json:"modified"`
	Status      string `json:"status"`
	UUID        string `json:"uuid"`
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
		Description: "New Timeslip",
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
		return fmt.Errorf("slip is already paused")
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

func (s *Slip) Resume() error {
	if s.Status == status.Started() {
		return fmt.Errorf("slip is already in progress")
	}

	s.Status = status.Started()
	s.Modified = timeNow()

	return nil
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

func (s Slip) TotalTimeWorked() int {
	if s.Status == status.Started() {
		return timeNow() - s.Modified + s.Worked
	}
	return s.Worked
}

func (s Slip) String() string {
	started := time.Unix(int64(s.Started), 0).Format("2006-01-02 15:04")
	timeWorked := s.DisplayTimeWorked(s.TotalTimeWorked())

	return fmt.Sprintf("%s | Started: %s | Worked: %s | Status: %s", s.FullName(), started, timeWorked, s.Status)
}

func (s Slip) DisplayTimeWorked(seconds int) string {
	var output string

	if seconds <= int(time.Duration(60*time.Second).Seconds()) {
		output = fmt.Sprintf("%d seconds", seconds)
	} else if seconds < int(time.Duration(60*time.Minute).Seconds()) {
		min := seconds / 60
		sec := seconds % 60
		if sec == 0 {
			output = fmt.Sprintf("%d minutes", min)
		} else {
			output = fmt.Sprintf("%dm %ds", min, sec)
		}
	} else {
		minutes := seconds / 60
		hour := minutes / 60
		min := minutes % 60
		if min == 0 {
			output = fmt.Sprintf("%d hours", hour)
		} else {
			output = fmt.Sprintf("%dh %dm", hour, min)
		}
	}

	return output
}

func (s Slip) ToJson() []byte {
	data, err := json.Marshal(s)
	if err != nil {
		return []byte{}
	}
	return data
}

func parseProjectName(name string) (string, string, error) {
	names := strings.Split(name, ".")

	switch len(names) {
	case 1:
		return names[0], "", nil
	case 2:
		return names[0], names[1], nil
	default:
		return "", "", fmt.Errorf("bad Project/Task name format. Expected 'ProjectName.TaskName' format")
	}
}

func timeNow() int {
	return int(time.Now().Unix())
}
