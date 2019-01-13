package reports

import (
	"bufio"
	"io"
	"sort"
	"strings"

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
// If a task name already exists then the time worked is added to the current
// entry, otherwise the timeslip is saved as a new entry. A new project is
// given the name as found in the first task.
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

	// add as a new task if its name has not been seen previously,
	// otherwise, add its worked time to the current task.
	if _, ok := p.tasks[t.name]; !ok {
		p.tasks[t.name] = t
	} else {
		p.tasks[t.name].timeWorked += t.timeWorked
	}

	p.totalTimeWorked += t.timeWorked

	return nil
}

// Check if the task was worked on during the desired time period.
//
// For simplicity, we only check that the finished time happened within the
// from/to time period, but what if 95% of the work was done outside of this?
// Perhaps a better solution would be to consider that at least 50% of the work
// was completed within the desired period.
func (p project) withinTimePeriod(t *task) bool {
	return t.finished >= int(p.timePeriod.From().Unix()) && t.finished <= int(p.timePeriod.To().Unix())
}

// Returns an array of all tasks, sorted by name.
func (p project) sortedTasks() []*task {
	var tasks []*task

	for _, t := range p.tasks {
		tasks = append(tasks, t)
	}

	sort.Slice(tasks, func(i, j int) bool {
		return strings.ToLower(tasks[i].name) < strings.ToLower(tasks[j].name)
	})

	return tasks
}
