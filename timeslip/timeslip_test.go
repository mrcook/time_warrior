package timeslip_test

import (
	"testing"

	"github.com/mrcook/time_warrior/timeslip"
)

func TestNewTimeSlipDefaults(t *testing.T) {
	ts, err := timeslip.New("timeWarrior.New")
	if err != nil {
		t.Errorf("Expected no error, got '%v'", err)
	}

	if ts.Project != "timeWarrior" {
		t.Errorf("Expected Project name 'timeWarrior', got '%s'", ts.Project)
	}
	if ts.Task != "New" {
		t.Errorf("Expected Task to be 'New', got '%s'", ts.Task)
	}
	if ts.Comment != "New Time Slip" {
		t.Errorf("Expected default comment to be 'New Time Slip', got '%s'", ts.Comment)
	}
	if ts.Started == 0 {
		t.Errorf("Expected time to have been set")
	}
	if ts.Modified != ts.Started {
		t.Errorf("Expected Modified time to Equal Started. Expected %d, got %d", ts.Started, ts.Modified)
	}
	if ts.Status != "started" {
		t.Errorf("Expected Status to be 'started', got '%s'", ts.Status)
	}
	if ts.UUID == "" || len(ts.UUID) != 36 {
		t.Errorf("Expected a valid UUID, got '%s'", ts.UUID)
	}
}

func TestOnlyProjectName(t *testing.T) {
	ts, err := timeslip.New("timeWarrior")
	if err != nil {
		t.Errorf("Expected no error, got '%v'", err)
	}

	if ts.Project != "timeWarrior" {
		t.Errorf("Expected Project to be 'timeWarrior', got '%s'", ts.Project)
	}
	if ts.Task != "" {
		t.Errorf("Expected Task to be blank, got '%s'", ts.Task)
	}
}

func TestProjectTaskName(t *testing.T) {
	ts, err := timeslip.New("timeWarrior.TaskName")
	if err != nil {
		t.Errorf("Expected no error, got '%v'", err)
	}

	if ts.Project != "timeWarrior" {
		t.Errorf("Expected Project to be 'timeWarrior', got '%s'", ts.Project)
	}
	if ts.Task != "TaskName" {
		t.Errorf("Expected Task to be 'TaskName', got '%s'", ts.Task)
	}
}

func TestMultiplePeriodsInProjectName(t *testing.T) {
	_, err := timeslip.New("timeWarrior.ProjectNames.BadName")
	if err == nil {
		t.Errorf("Expected an Error due to multiple periods being used.")
	}

	errorMessage := "Bad Project/Task name format. Expected 'ProjectName.TaskName' format"
	if err.Error() != errorMessage {
		t.Errorf("Expected '%s' error, got '%v'", errorMessage, err)
	}
}
