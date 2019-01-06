// Package timeslip creates a new timeslip, along with managing its state.
package timeslip

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/mrcook/time_warrior/timeslip/status"
	"github.com/mrcook/time_warrior/timeslip/worked"
)

// Slip represents a timeslip.
// Note: timestamps are stored as Unix time.
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

// New returns a new "started" timeslip.
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

// Unmarshal parses the JSON-encoded data and stores the result
// in the struct pointed to by `slip`.
func Unmarshal(data []byte, slip *Slip) error {
	return json.Unmarshal(data, slip)
}

// Pause a started timeslip.
func (s *Slip) Pause() error {
	if s.isPaused() {
		return fmt.Errorf("slip is already paused")
	}

	s.Status = status.Paused()
	s.Worked += timeNow() - s.Modified
	s.Modified = timeNow()

	return nil
}

func (s Slip) isPaused() bool {
	return s.Status == status.Paused()
}

// Resume a paused timeslip.
func (s *Slip) Resume() error {
	if s.Status == status.Started() {
		return fmt.Errorf("slip is already in progress")
	}

	s.Status = status.Started()
	s.Modified = timeNow()

	return nil
}

// Done marks a timeslip as completed.
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

// Adjust the current worked time from the given string value.
// Note: the modified time should be moved forward with the
// adjustment, unless that would put it into the future, in
// which case the started time should be pushed back.
func (s *Slip) Adjust(adjustment string) error {
	if !s.isPaused() {
		return fmt.Errorf("only paused timeslips can be changed")
	}

	a := worked.WorkTime{}
	if err := a.FromString(adjustment); err != nil {
		return err
	}

	w := worked.WorkTime{}
	w.FromSeconds(s.Worked)

	w.Add(&a)
	s.Worked = w.ToSeconds()

	if s.Worked < 0 {
		s.Worked = 0
		return nil
	}

	now := timeNow()
	workedTime := s.Started + s.Worked

	if workedTime > s.Modified && workedTime <= now {
		s.Modified = workedTime
	} else if workedTime > now {
		s.Started = s.Modified - s.Worked
	}

	return nil
}

// Name returns the full timeslip name as `Project.Task`, or just `Project` if no task is present.
func (s Slip) Name() string {
	if s.Task == "" {
		return s.Project
	}
	return s.Project + "." + s.Task
}

// TotalTimeWorked returns the total time worked on a timeslip.
// If a timeslip is currently started, the worked time is adjusted based
// on the modified and current time.
func (s Slip) TotalTimeWorked() int {
	if s.Status == status.Started() {
		return timeNow() - s.Modified + s.Worked
	}
	return s.Worked
}

// String returns a CLI friendly representation of the timeslip.
func (s Slip) String() string {
	started := time.Unix(int64(s.Started), 0).Format("2006-01-02 15:04")

	w := worked.WorkTime{}
	w.FromSeconds(s.TotalTimeWorked())

	return fmt.Sprintf("%s | Started: %s | Worked: %s | Status: %s", s.Name(), started, w.String(), s.Status)
}

// ToJson converts a timeslip to a JSON string.
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

// timeNow returns a Unix timestamp.
func timeNow() int {
	return int(time.Now().Unix())
}
