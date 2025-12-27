package timeslip_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/mrcook/time_warrior/timeslip"
	"github.com/mrcook/time_warrior/timeslip/status"
)

// TODO: some of these tests will fail when the second ticks over!

func TestSlip_New(t *testing.T) {
	t.Run("with both project and task name", func(t *testing.T) {
		ts, err := timeslip.New("timeWarrior.TaskName")
		if err != nil {
			t.Errorf("unexpected error, got '%s'", err)
		}

		if ts.Project != "timeWarrior" {
			t.Errorf("expected project name to be set, got '%s'", ts.Project)
		}

		if ts.Task != "TaskName" {
			t.Errorf("expected task name to be set, got '%s'", ts.Task)
		}
	})

	t.Run("when only project name is given", func(t *testing.T) {
		ts, err := timeslip.New("timeWarrior")
		if err != nil {
			t.Errorf("unexpected error, got '%s'", err)
		}

		if ts.Project != "timeWarrior" {
			t.Errorf("expected project name to be set, got '%s'", ts.Project)
		}

		if ts.Task != "" {
			t.Errorf("expected task name to be blank, got '%s'", ts.Task)
		}
	})

	t.Run("when multiple periods included", func(t *testing.T) {
		_, err := timeslip.New("timeWarrior.ProjectNames.BadName")
		if err == nil {
			t.Fatalf("expected an error")
		}

		errorMessage := "bad Project/Task name format. Expected 'ProjectName.TaskName' format"
		if err.Error() != errorMessage {
			t.Errorf("expected '%s', got '%s'", errorMessage, err)
		}
	})
}

func TestSlip_New_withDefaults(t *testing.T) {
	ts, err := timeslip.New("timeWarrior.New")
	if err != nil {
		t.Errorf("unxpected error, got '%s'", err)
	}

	if ts.Project != "timeWarrior" {
		t.Errorf("expected project name to be set, got '%s'", ts.Project)
	}

	if ts.Task != "New" {
		t.Errorf("expected task name to be set, got '%s'", ts.Task)
	}

	if ts.Description != "New Timeslip" {
		t.Errorf("expected default comment, got '%s'", ts.Description)
	}

	if ts.Started == 0 {
		t.Error("expected start time to have been set")
	}

	if ts.Modified != ts.Started {
		t.Errorf("expected modified time %d to equal started time %d", ts.Modified, ts.Started)
	}

	if ts.Status != status.Started {
		t.Errorf("expected '%s' status, got '%s'", status.Started, ts.Status)
	}

	if ts.UUID == "" {
		t.Error("expected an UUID to be set")
	}
}

func TestSlip_UUIDGeneration(t *testing.T) {
	ts, _ := timeslip.New("NewUUID")

	if len(ts.UUID) != 36 {
		t.Errorf("expected UUID with 36 characters, got %d", len(ts.UUID))
	}

	_, err := uuid.Parse(ts.UUID)
	if err != nil {
		t.Errorf("unexpected UUID error, got %s", err)
	}
}

func TestSlip_Pause(t *testing.T) {
	t.Run("pause a timeslip", func(t *testing.T) {
		unixNow := int(time.Now().Unix())
		ts := timeslip.Slip{
			Started:  unixNow - 60,
			Worked:   0,
			Modified: unixNow - 60,
			Status:   status.Started,
		}
		modifiedWas := ts.Modified

		err := ts.Pause()
		if err != nil {
			t.Fatalf("unexpected error on pause %s", err)
		}

		if ts.Status != status.Paused {
			t.Errorf("expected '%s' status, got '%s'", status.Paused, ts.Status)
		}

		if ts.Modified == modifiedWas {
			t.Errorf("expected modified time %d to have been updated", ts.Modified)
		}

		if ts.Worked != 60 {
			t.Errorf("expected worked time to have been updated, got %d", ts.Worked)
		}
	})

	t.Run("pausing an already paused timeslip", func(t *testing.T) {
		ts, _ := timeslip.New("Project.NotRePausable")
		_ = ts.Pause()

		if err := ts.Pause(); err == nil {
			t.Errorf("expected error when pausing a paused slip")
		}
	})
}

func TestSlip_Resume(t *testing.T) {
	t.Run("resume a paused timeslip", func(t *testing.T) {
		unixNow := int(time.Now().Unix())
		ts := timeslip.Slip{
			Started:  unixNow - 60,
			Worked:   30,
			Modified: unixNow - 30,
			Status:   status.Paused,
		}
		modifiedWas := ts.Modified

		err := ts.Resume()
		if err != nil {
			t.Fatalf("unexpected resume error, got %s", err.Error())
		}

		if ts.Status != status.Resumed {
			t.Errorf("expected '%s' status, got '%s'", status.Resumed, ts.Status)
		}

		if ts.Modified == modifiedWas {
			t.Errorf("modified time was not updated")
		}

		if ts.Worked != 30 {
			t.Errorf("worked time should remain unchanged, got %d", ts.Worked)
		}
	})

	t.Run("resuming a started timeslip", func(t *testing.T) {
		ts, _ := timeslip.New("Resume.Started")

		err := ts.Resume()
		if !strings.Contains(err.Error(), "slip is already in progress") {
			t.Errorf("unexpected resume error, got %s", err.Error())
		}
	})

	t.Run("resuming a resumed timeslip", func(t *testing.T) {
		ts, _ := timeslip.New("Resume.Resumed")
		_ = ts.Pause()
		_ = ts.Resume()

		err := ts.Resume()
		if !strings.Contains(err.Error(), "slip is already in progress") {
			t.Errorf("unexpected resume error, got %s", err.Error())
		}
	})
}

func TestSlip_Done(t *testing.T) {
	unixNow := int(time.Now().Unix())

	t.Run("finished a started timeslip", func(t *testing.T) {
		ts := timeslip.Slip{
			Started:  unixNow - 60,
			Worked:   0,
			Modified: unixNow - 60,
			Status:   status.Started,
		}
		modifiedWas := ts.Modified

		description := "Write tests for completing a started timeslip"
		ts.Done(description)

		if ts.Status != status.Completed {
			t.Errorf("unexpected status got '%s'", ts.Status)
		}

		if ts.Description != description {
			t.Errorf("unexpected description got '%s'", ts.Description)
		}

		if ts.Worked != 60 {
			t.Errorf("unexpected worked time, got %d", ts.Worked)
		}

		if ts.Finished == 0 {
			t.Error("expected finished time to have been updated")
		}

		if ts.Modified == modifiedWas {
			t.Error("expected modified time to have been updated")
		}
	})

	t.Run("finishing a paused timeslip", func(t *testing.T) {
		ts := timeslip.Slip{
			Started:  unixNow - 60,
			Worked:   30,
			Modified: unixNow - 30,
			Status:   status.Paused,
		}
		modifiedWas := ts.Modified

		ts.Done("Write tests for completing a paused timeslip")

		if ts.Finished == 0 {
			t.Error("expected finished time to have been updated")
		}

		if ts.Finished > modifiedWas {
			t.Errorf("expected finished time %d to equal modified time %d", ts.Finished, modifiedWas)
		}

		if ts.Worked != 30 {
			t.Errorf("worked time should not have changed, got %d", ts.Worked)
		}

		if ts.Modified != modifiedWas {
			t.Error("modified time should not have changed")
		}
	})
}

func TestSlip_NewFromJSONUnmarshallError(t *testing.T) {
	var badJSON = []byte(`"project": "BadWarrior", "task":`)

	slip := &timeslip.Slip{}
	err := timeslip.Unmarshal(badJSON, slip)
	if err == nil {
		t.Error("expected an unmarshal error, got none")
	}
}

func TestSlip_NewFromJSON(t *testing.T) {
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

	slip := &timeslip.Slip{}
	err := timeslip.Unmarshal(jsonBlob, slip)
	if err != nil {
		t.Error("failed to unmarshal the JSON data: ", err)
	}

	if slip.Project != "timeWarrior" {
		t.Errorf("expected project name, got '%s'", slip.Project)
	}
	if slip.Task != "NewFromJSON" {
		t.Errorf("expected task name, got '%s'", slip.Task)
	}
	if slip.Description != "Testing import of JSON" {
		t.Errorf("expected description, got '%s'", slip.Description)
	}
	if slip.Started != 1442669540 {
		t.Errorf("expected started time, got '%d'", slip.Started)
	}
	if slip.Worked != 472 {
		t.Errorf("expected time worked, got '%d'", slip.Worked)
	}
	if slip.Finished != 1442674736 {
		t.Errorf("expected finished time, got '%d'", slip.Finished)
	}
	if slip.Modified != 1442674736 {
		t.Errorf("expected modified time, got '%d'", slip.Modified)
	}
	if slip.Status != status.Completed {
		t.Errorf("expected '%s' status, got '%s'", status.Completed, slip.Status)
	}
	if slip.UUID != "0d8e895e-d3db-4887-86e3-8bb7f63ba101" {
		t.Errorf("expected UUID, got '%s'", slip.UUID)
	}
}

func TestSlip_Name(t *testing.T) {
	t.Run("when both project and task names are present", func(t *testing.T) {
		slip := timeslip.Slip{Project: "TimeWarrior", Task: "Name"}

		if slip.Name() != "TimeWarrior.Name" {
			t.Errorf("expected Project and Task names joined with a period, got: %s", slip.Name())
		}
	})

	t.Run("when only project name is present", func(t *testing.T) {
		slip := timeslip.Slip{Project: "TimeWarrior", Task: ""}

		if slip.Name() != "TimeWarrior" {
			t.Errorf("expected the project name, got '%s'", slip.Name())
		}
	})
}

func TestSlip_StringOutput(t *testing.T) {
	now := time.Now()

	t.Run("finished timeslip after it was paused", func(t *testing.T) {
		started := 1706833230
		worked := 10
		modified := started + worked
		finished := modified + 10

		ts := timeslip.Slip{
			Project:  "timeWarrior",
			Task:     "DoneAfterPause",
			Started:  started,
			Worked:   worked,
			Finished: finished,
			Modified: modified,
			Status:   status.Completed,
		}
		expectedOutput := "timeWarrior.DoneAfterPause | Started: 2024-02-02 01:20 | Worked: 10 seconds | Status: completed"

		output := ts.String()
		if output != expectedOutput {
			t.Errorf("formatting incorrect:\n     got: %s\nexpected: %s", output, expectedOutput)
		}
	})

	t.Run("it should output the correct format", func(t *testing.T) {
		sixMinutes := 6 * time.Minute
		startedAgo := now.Add(-sixMinutes)

		ts := timeslip.Slip{
			Project:  "timeWarrior",
			Task:     "String",
			Started:  int(startedAgo.Unix()),
			Worked:   int(sixMinutes.Seconds()),
			Modified: int(now.Unix()),
			Status:   status.Paused,
		}
		startStr := startedAgo.Format("2006-01-02 15:04")
		modifiedStr := now.Format("2006-01-02 15:04")

		expectedOutput := fmt.Sprintf("timeWarrior.String | Started: %s | Worked: 6 minutes | Status: paused (%s)", startStr, modifiedStr)

		output := ts.String()
		if output != expectedOutput {
			t.Errorf("formatting incorrect:\n     got: %s\nexpected: %s", output, expectedOutput)
		}
	})
}

func TestSlip_TotalTimeWorked(t *testing.T) {
	tests := map[string]struct {
		status                    string
		started, worked, modified int // in seconds ago
		expected                  int
	}{
		"started, with no pauses/resumes": {
			status.Started, 100, 0, 100, 100,
		},
		"paused, with no previous pauses": {
			status.Paused, 120, 30, 120 - 30, 30,
		},
		"paused, with multiple previous pauses": {
			status.Paused, 800, 5, 100, 5,
		},
		"resumed after a pause": {
			status.Resumed, 60, 22, 7, 22 + 7,
		},
		"completed without any pauses": {
			status.Completed, 488, 488, 0, 488,
		},
		"completed after a pause": {
			status.Completed, 65535, 30, 0, 30,
		},
	}

	now := time.Now()
	for msg, test := range tests {
		t.Run(msg, func(t *testing.T) {
			ts := timeslip.Slip{
				Started:  int(now.Add(-time.Duration(test.started) * time.Second).Unix()),
				Worked:   test.worked,
				Modified: int(now.Add(-time.Duration(test.modified) * time.Second).Unix()),
				Status:   test.status,
			}

			total := ts.TotalTimeWorked()
			if total != test.expected {
				t.Errorf("expected %d seconds worked, got %d", test.expected, total)
			}
		})
	}
}

func TestSlip_ToJson(t *testing.T) {
	slip, _ := timeslip.New("Output.ToJson")
	expectedOutput := fmt.Sprintf(`{"project":"Output","task":"ToJson","description":"New Timeslip","started":%d,"worked":0,"finished":0,"modified":%d,"status":"started","uuid":"%s"}`, slip.Started, slip.Modified, slip.UUID)

	output := slip.ToJson()
	if string(output) != expectedOutput {
		t.Errorf("formatting incorrect:\n     got: %s\nexpected: %s", output, expectedOutput)
	}
}

func TestSlip_AdjustWorked(t *testing.T) {
	newTimeslip := func(started, modified, worked int, state string) timeslip.Slip {
		unixNow := int(time.Now().Unix())

		return timeslip.Slip{
			Project:  "timeWarrior",
			Task:     "Adjust",
			Started:  unixNow - started,
			Worked:   worked,
			Modified: unixNow - modified,
			Status:   state,
		}
	}

	t.Run("with enough time to adjust without affecting started/modified times", func(t *testing.T) {
		adjustment := 10

		slip := newTimeslip(60, 30, 10, status.Paused)

		startedWas := slip.Started
		modifiedWas := slip.Modified

		err := slip.Adjust(fmt.Sprintf("%ds", adjustment))
		if err != nil {
			t.Fatalf("unexpected error on adjust %s", err)
		}

		// worked + adjustment
		if slip.Worked != 10+adjustment {
			t.Errorf("expected worked time to have been incremented, got %d", slip.Worked)
		}

		if slip.Started != startedWas {
			t.Errorf("expected started time %d to remain unchanged, got %d", startedWas, slip.Started)
		}

		if slip.Modified != modifiedWas {
			t.Errorf("expected modified time %d to remain unchanged, got %d", modifiedWas, slip.Modified)
		}
	})

	t.Run("with enough time since last pause, update the modified time", func(t *testing.T) {
		adjustment := 15

		slip := newTimeslip(60, 30, 30, status.Paused)
		startedWas := slip.Started
		modifiedWas := slip.Modified

		err := slip.Adjust(fmt.Sprintf("%ds", adjustment))
		if err != nil {
			t.Fatalf("unexpected error on adjust %s", err)
		}

		if slip.Worked != 30+adjustment {
			t.Errorf("expected worked time to have been incremented, got %d", slip.Worked)
		}

		if slip.Started != startedWas {
			t.Errorf("expected started time %d to remain unchanged, got %d", startedWas, slip.Started)
		}

		if slip.Modified != modifiedWas+adjustment {
			t.Errorf("expected modified time to have been moved forward to %d, got %d", modifiedWas+adjustment, slip.Modified)
		}
	})

	t.Run("when not enough modified time, start time is adjusted", func(t *testing.T) {
		adjustment := 15 // would make modified 5 seconds in the future

		slip := newTimeslip(60, 10, 50, status.Paused)
		startedWas := slip.Started
		modifiedWas := slip.Modified

		err := slip.Adjust(fmt.Sprintf("%ds", adjustment))
		if err != nil {
			t.Fatalf("unexpected error on adjust %s", err)
		}

		if slip.Worked != 50+adjustment {
			t.Errorf("expected worked time to have been incremented, got %d", slip.Worked)
		}

		if slip.Modified != modifiedWas {
			t.Errorf("expected modified time %d to remain unchanged, got %d", modifiedWas, slip.Modified)
		}

		newStartTime := startedWas - adjustment
		if slip.Started != newStartTime {
			t.Errorf("expected started time to have been pushed back to %d, got %d", newStartTime, slip.Started)
		}
	})

	t.Run("negative adjustments", func(t *testing.T) {
		t.Run("should only adjust the worked time", func(t *testing.T) {
			adjustment := -10

			slip := newTimeslip(60, 10, 50, status.Paused)
			startedWas := slip.Started
			modifiedWas := slip.Modified

			err := slip.Adjust(fmt.Sprintf("%ds", adjustment))
			if err != nil {
				t.Fatalf("unexpected error on adjust %s", err)
			}

			if slip.Worked != 40 {
				t.Errorf("expected worked time to be decremented to %d, got %d", 40, slip.Worked)
			}

			if slip.Started != startedWas {
				t.Errorf("expected started time %d to be unchanged, got %d", startedWas, slip.Started)
			}

			if slip.Modified != modifiedWas {
				t.Errorf("expected modified time %d to be unchanged, got %d", modifiedWas, slip.Modified)
			}
		})

		t.Run("when adjustment is more than worked time", func(t *testing.T) {
			adjustment := -20

			slip := newTimeslip(60, 50, 10, status.Paused)

			err := slip.Adjust(fmt.Sprintf("%ds", adjustment))
			if err != nil {
				t.Fatalf("unexpected error on adjust %s", err)
			}

			if slip.Worked != 0 {
				t.Errorf("worked time should be set to zero, got %d", slip.Worked)
			}
		})
	})

	t.Run("when adjustment is missing the time unit", func(t *testing.T) {
		slip := newTimeslip(60, 0, 60, status.Paused)

		worked := slip.Worked

		err := slip.Adjust("1")
		if err == nil {
			t.Fatalf("expected an error")
		}

		if err.Error() != "invalid time unit, got '1'" {
			t.Errorf("expected invalid unit error, got '%s'", err)
		}

		if slip.Worked != worked {
			t.Errorf("expected worked time %d to be unchanged, got %d", worked, slip.Worked)
		}
	})

	t.Run("with a started timeslip", func(t *testing.T) {
		adjustment := 10

		slip := newTimeslip(10, 10, 0, status.Started)
		startedWas := slip.Started

		err := slip.Adjust(fmt.Sprintf("%ds", adjustment))
		if err != nil {
			t.Fatalf("unexpected error on adjust %s", err)
		}

		// 10 since starting + adjustment of 10
		if slip.Worked != 20 {
			t.Errorf("expected worked time to have been incremented, got %d", slip.Worked)
		}

		newStarted := startedWas - adjustment // move back started for adjustment
		if slip.Started != newStarted {
			t.Errorf("expected started time %d to be moved back, got %d", newStarted, slip.Started)
		}

		newModified := int(time.Now().Unix()) // because Adjust() pauses the timeslip
		if slip.Modified != newModified {
			t.Errorf("expected modified time %d to be current time, got %d", newModified, slip.Modified)
		}

		if slip.Status != status.Started {
			t.Errorf("expected status to remain unchanged, got '%s'", slip.Status)
		}
	})

	t.Run("when timeslip has been resumed", func(t *testing.T) {
		adjustment := 300

		slip := newTimeslip(60, 5, 55, status.Resumed)
		startedWas := slip.Started

		err := slip.Adjust(fmt.Sprintf("%ds", adjustment))
		if err != nil {
			t.Fatalf("unexpected error on adjust %s", err)
		}

		// 55 + 5 since resuming + adjustment of 300
		if slip.Worked != 360 {
			t.Errorf("expected worked time to have been incremented, got %d", slip.Worked)
		}

		newStarted := startedWas - adjustment // move back started for adjustment
		if slip.Started != newStarted {
			t.Errorf("expected started time %d to be moved back, got %d", newStarted, slip.Started)
		}

		newModified := int(time.Now().Unix()) // because Adjust() pauses the timeslip
		if slip.Modified != newModified {
			t.Errorf("expected modified time %d to be current time, got %d", newModified, slip.Modified)
		}

		if slip.Status != status.Resumed {
			t.Errorf("expected status to remain unchanged, got '%s'", slip.Status)
		}
	})
}
