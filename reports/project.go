package reports

import (
	"bufio"
	"fmt"
	"io"
	"sort"

	"github.com/mrcook/time_warrior/reports/period"
)

type project struct {
	name            string
	timePeriod      *period.Period
	totalTimeWorked int
	tasks           map[string]*task
	scanErrors      []scanError
}

type scanError struct {
	scanner  error
	timeslip string
}

// Initializes a new project.
func newProject(p *period.Period) *project {
	return &project{
		timePeriod: p,
		tasks:      make(map[string]*task),
	}
}

// Process the timeslips form the given file. Scanner errors are returned
// directly, project/task errors are recorded for later use.
func (p *project) process(file io.Reader) error {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		err := p.processSlip(scanner.Bytes())
		if err != nil {
			p.scanErrors = append(p.scanErrors, scanError{scanner: err, timeslip: scanner.Text()})
		}
	}
	return scanner.Err()
}

// Process a timeslip as a task and add the time worked to the project total.
//
// Each timeslip is saved as a new task entry, even if the task name already exists.
// This preserves the individual start and finish times for each task instance.
func (p *project) processSlip(data []byte) error {
	t, err := newTask(data)
	if err != nil {
		return err
	}

	// if this is a new project, we set its name
	if p.name == "" {
		p.name = t.project
	}

	// skip processing if the task was done outside the desired time period
	if p.timePeriod.IsSet() && !p.withinTimePeriod(t) {
		return nil
	}

	// Create a unique key for this task instance
	taskKey := fmt.Sprintf("%s_%d", t.name, t.started)
	p.tasks[taskKey] = t
	p.totalTimeWorked += t.timeWorked

	return nil
}

// Check if the task was worked on during the desired time period.
//
// A task is considered within the time period if either:
// 1. It started within the period
// 2. It finished within the period
// 3. It spans across the period (started before and finished after)
func (p project) withinTimePeriod(t *task) bool {
	periodStart := int(p.timePeriod.From().Unix())
	periodEnd := int(p.timePeriod.To().Unix())

	// Task started within period
	startedInPeriod := t.started >= periodStart && t.started <= periodEnd

	// Task finished within period
	finishedInPeriod := t.finished >= periodStart && t.finished <= periodEnd

	// Task spans across period
	spansPeriod := t.started <= periodStart && (t.finished == 0 || t.finished >= periodEnd)

	return startedInPeriod || finishedInPeriod || spansPeriod
}

// Returns an array of all tasks, sorted by start time.
func (p project) sortedTasks() []*task {
	var tasks []*task

	for _, t := range p.tasks {
		tasks = append(tasks, t)
	}

	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].started < tasks[j].started
	})

	return tasks
}

// Name returns the project name
func (p *project) Name() string {
	return p.name
}

// SortedTasks returns a sorted list of tasks
func (p *project) SortedTasks() []*task {
	return p.sortedTasks()
}
