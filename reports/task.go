package reports

import "github.com/mrcook/time_warrior/timeslip"

type task struct {
	name       string
	project    string
	started    int
	finished   int
	timeWorked int
}

// Creates a new task from a timeslip JSON string.
func newTask(jsonData []byte) (*task, error) {
	slip := &timeslip.Slip{}
	if err := timeslip.Unmarshal(jsonData, slip); err != nil {
		return nil, err
	}

	var name string
	if slip.Task == "" {
		name = "."
	} else {
		name = slip.Task
	}

	t := &task{
		name:       name,
		project:    slip.Project,
		started:    slip.Started,
		finished:   slip.Finished,
		timeWorked: slip.Worked,
	}

	return t, nil
}

// Name returns the task name
func (t *task) Name() string {
	return t.name
}

// Started returns the task start time
func (t *task) Started() int {
	return t.started
}

// Finished returns the task finish time
func (t *task) Finished() int {
	return t.finished
}
