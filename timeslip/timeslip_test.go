package timeslip_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/mrcook/time_warrior/timeslip"
	"github.com/mrcook/time_warrior/timeslip/status"
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
		t.Error("Expected time to have been set")
	}
	if ts.Modified != ts.Started {
		t.Errorf("Expected Modified time to Equal Started. Expected %d, got %d", ts.Started, ts.Modified)
	}
	if ts.Status != status.Started() {
		t.Errorf("Expected Status to be '%s', got '%s'", status.Started(), ts.Status)
	}
	if ts.UUID == "" {
		t.Error("Expected an UUID to be set")
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

func TestUUIDGeneration(t *testing.T) {
	ts, _ := timeslip.New("NewUUID")

	if len(ts.UUID) != 36 {
		t.Error("Expected the UUID to be 36 characters long")
	}

	_, err := uuid.Parse(ts.UUID)
	if err != nil {
		t.Errorf("Expected a valid UUID, got error: %v", err)
	}
}

func TestPausing(t *testing.T) {
	ts, _ := timeslip.New("Project.Pause")
	ts.Started -= 100
	ts.Modified -= 100

	worked := ts.Worked
	modified := ts.Modified

	ts.Pause()

	if ts.Status != status.Paused() {
		t.Errorf("Expected status to be '%s', got '%s'", status.Paused(), ts.Status)
	}
	if ts.Modified < modified+100 {
		t.Errorf("Expected updated Modified date, original: %d, current: %d", modified, ts.Modified)
	}
	if ts.Worked < 100 {
		t.Errorf("Expected updated worked time, got %d, was %d", ts.Worked, worked)
	}
}

func TestPausingAlreadyPausedSlip(t *testing.T) {
	ts, _ := timeslip.New("Project.PausePaused")
	ts.Started -= 100
	ts.Modified -= 100
	ts.Pause()

	if err := ts.Pause(); err == nil {
		t.Errorf("Expected an error when pausing a paused slip")
	}
}

func TestResuming(t *testing.T) {
	ts, _ := timeslip.New("Project.Resume")
	ts.Started -= 200
	ts.Modified -= 200
	ts.Status = status.Paused()

	ts.Resume()

	if ts.Status != status.Started() {
		t.Errorf("Expected status to be '%s', got '%s'", status.Started(), ts.Status)
	}
	currentTime := int(time.Now().Unix())
	if ts.Modified < currentTime-2 {
		t.Errorf("Expected Modified to be updated, current time is %d, got %d", currentTime, ts.Modified)
	}
}
