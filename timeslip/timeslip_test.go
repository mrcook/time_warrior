package timeslip_test

import (
	"fmt"
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
	if ts.Description != "New Time Slip" {
		t.Errorf("Expected default comment to be 'New Time Slip', got '%s'", ts.Description)
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

func TestFinishingAStartedTask(t *testing.T) {
	tenMinutesInSeconds := int((10 * time.Minute).Seconds())
	description := "Write tests for completing a started time slip"

	ts, _ := timeslip.New("Work.Done")
	ts.Started -= tenMinutesInSeconds
	ts.Modified -= tenMinutesInSeconds

	modifiedTime := ts.Modified

	ts.Done(description)

	if ts.Status != status.Completed() {
		t.Errorf("Expected status to be '%s', got '%s'", status.Completed(), ts.Status)
	}
	if ts.Description != description {
		t.Errorf("Expected description to be '%s', got '%s'", description, ts.Description)
	}
	if ts.Worked < tenMinutesInSeconds {
		t.Errorf("Expected worked time to be at least %d seconds, got %d", tenMinutesInSeconds, ts.Worked)
	}
	if ts.Finished == 0 {
		t.Errorf("Expected finished time to have been updated, got %d", ts.Finished)
	}
	if ts.Modified == modifiedTime {
		t.Error("Expected modified time to have been updated")
	}
}

func TestFinishingAPausedTask(t *testing.T) {
	oneHourAgo := int(time.Now().Add(-1 * time.Hour).Unix())
	halfHourAgo := int(time.Now().Add(-30 * time.Minute).Unix())
	fifteenMinutes := int((15 * time.Minute).Seconds())

	ts := timeslip.Slip{
		Started:  oneHourAgo,
		Worked:   fifteenMinutes,
		Modified: halfHourAgo,
		Status:   status.Paused(),
	}
	modifiedTime := halfHourAgo

	ts.Done("Write tests for completing a paused time slip")

	if ts.Finished > modifiedTime {
		t.Errorf("Expected finished time '%d', to equal original modified time '%d'", ts.Finished, modifiedTime)
	}
	if ts.Worked != fifteenMinutes {
		t.Errorf("Expected time worked to not change from %d, got %d", fifteenMinutes, ts.Worked)
	}
	if ts.Modified == modifiedTime {
		t.Error("Expected modified time to have been updated")
	}
}

func TestNewFromJSONUnmarshallError(t *testing.T) {
	var badJSON = []byte(`"project": "BadWarrior", "task":`)

	_, err := timeslip.NewFromJSON(badJSON)
	if err == nil {
		t.Error("Expected unamrshal error")
	}
}

func TestNewFromJSON(t *testing.T) {
	var jsonBlob = []byte(`{
		"project": "timeWarrior",
		"task": "NewFromJSON",
		"description": "Testing import of JSON",
		"started": 1442669540,
		"worked": 472,
		"finished": 1442674736,
		"modified": 1442674736,
		"status": "completed",
		"uuid": "0d8e895e-d3db-4887-86e3-8bb7f63ba101"
	}`)

	slip, err := timeslip.NewFromJSON(jsonBlob)
	if err != nil {
		t.Error("Failed to unamrshal the JSON data: ", err)
	}

	if slip.Project != "timeWarrior" {
		t.Errorf("Expected project name, got '%s'", slip.Project)
	}
	if slip.Task != "NewFromJSON" {
		t.Errorf("Expected task name, got '%s'", slip.Task)
	}
	if slip.Description != "Testing import of JSON" {
		t.Errorf("Expected description, got '%s'", slip.Description)
	}
	if slip.Started != 1442669540 {
		t.Errorf("Expected started time, got '%d'", slip.Started)
	}
	if slip.Worked != 472 {
		t.Errorf("Expected time worked, got '%d'", slip.Worked)
	}
	if slip.Finished != 1442674736 {
		t.Errorf("Expected finished time, got '%d'", slip.Finished)
	}
	if slip.Modified != 1442674736 {
		t.Errorf("Expected modified time, got '%d'", slip.Modified)
	}
	if slip.Status != status.Completed() {
		t.Errorf("Expected status to be '%s', got '%s'", status.Completed(), slip.Status)
	}
	if slip.UUID != "0d8e895e-d3db-4887-86e3-8bb7f63ba101" {
		t.Errorf("Expected UUID, got '%s'", slip.UUID)
	}
}

func TestFullName(t *testing.T) {
	slip := timeslip.Slip{Project: "timeWarrior", Task: "FullName"}

	if slip.FullName() != "timeWarrior.FullName" {
		t.Errorf("Expected Project and Task names joined with a perdiod, got: %s", slip.FullName())
	}

	slip.Task = ""
	if slip.FullName() != "timeWarrior" {
		t.Errorf("Expected just the project name, got: %s", slip.FullName())
	}
}

func TestStringOutput(t *testing.T) {
	sixMinutesAgo := time.Now().Add(-6 * time.Minute)
	started := int(sixMinutesAgo.Unix())
	formattedStart := sixMinutesAgo.Format("2006-01-02 15:04")

	slip, _ := timeslip.New("timeWarrior.String")
	slip.Started = started
	slip.Modified = started
	slip.Pause()

	expectedOutput := fmt.Sprintf("timeWarrior.String | Started: %s | Worked: 6 mins | Status: paused", formattedStart)

	output := slip.String()
	if output != expectedOutput {
		t.Errorf("Formatting incorrect:\n     got: %s\nexpected: %s", output, expectedOutput)
	}
}
