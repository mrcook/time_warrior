package reports

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/mrcook/time_warrior/reports/period"
	"github.com/mrcook/time_warrior/timeslip"
	"github.com/mrcook/time_warrior/timeslip/worked"
)

type Report struct {
	PendingTimeslip timeslip.Slip

	timePeriod      *period.Period
	totalTimeWorked int
	projects        []*project
	errors          []error
}

// Returns a new report using the given time unit.
func New(timeUnit string) *Report {
	return &Report{timePeriod: period.Parse(timeUnit)}
}

// ProcessProjectFile reads a project file and processes all of its timeslips,
// calculating the time worked for each task.
func (r *Report) ProcessProjectFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		// it's okay to stop the program for file read errors
		log.Fatal(err)
	}
	defer file.Close()

	p := newProject(r.timePeriod)
	if err := p.process(file); err != nil {
		r.errors = append(r.errors, err)
	}

	r.totalTimeWorked += p.totalTimeWorked
	r.projects = append(r.projects, p)
}

// PrintReport prints a report to the terminal for the projects/tasks.
func (r Report) PrintReport() {
	if len(r.projects) == 1 {
		r.printProjectTasks()
	} else if len(r.projects) > 1 {
		r.sortProjectsByName()
		r.printProjects()
	} else {
		fmt.Println("No available data.")
	}

	if len(r.errors) > 0 {
		fmt.Println()
		r.printErrors()
	}
}

// Displays all projects with their time worked to the terminal.
func (r Report) printProjects() {
	if r.timePeriod.IsSet() {
		fmt.Printf("Time Period: %s (%s)\n\n", r.timePeriod.Period(), r.formattedDates())
	}

	for _, p := range r.projects {
		if p.totalTimeWorked == 0 {
			continue
		}

		w := worked.WorkTime{}
		w.FromSeconds(p.totalTimeWorked)
		if w.Hours == 0 {
			fmt.Printf("     %4dm : %s\n", w.Minutes, p.name)
		} else {
			fmt.Printf("%4dh %3dm : %s\n", w.Hours, w.Minutes, p.name)
		}
	}

	if r.PendingTimeslip.TotalTimeWorked() > 0 {
		fmt.Println("-----------")

		pending := worked.WorkTime{}
		pending.FromSeconds(r.PendingTimeslip.TotalTimeWorked())
		fmt.Printf("%4dh %3dm : %s pending timeslip\n", pending.Hours, pending.Minutes, r.PendingTimeslip.Project)
	}

	r.printTotal(r.totalTimeWorked + r.PendingTimeslip.TotalTimeWorked())
}

// Displays project overview, along with all tasks and their time worked.
func (r Report) printProjectTasks() {
	p := r.projects[0]

	fmt.Printf("Project Name: %s\n", p.name)
	if r.timePeriod.IsSet() {
		fmt.Printf("Time Period:  %s (%s)\n", r.timePeriod.Period(), r.formattedDates())
	}
	fmt.Println()

	fmt.Println("Task List")

	for _, t := range p.sortedTasks() {
		w := worked.WorkTime{}
		w.FromSeconds(t.timeWorked)
		if w.Hours == 0 {
			fmt.Printf("     %4dm : %s\n", w.Minutes, t.name)
		} else {
			fmt.Printf("%4dh %3dm : %s\n", w.Hours, w.Minutes, t.name)
		}
	}

	if p.name == r.PendingTimeslip.Project && r.PendingTimeslip.TotalTimeWorked() > 0 {
		fmt.Println("-----------")

		pending := worked.WorkTime{}
		pending.FromSeconds(r.PendingTimeslip.TotalTimeWorked())

		task := ""
		if r.PendingTimeslip.Task != "" {
			task = fmt.Sprintf("%s ", r.PendingTimeslip.Task)
		}

		fmt.Printf("%4dh %3dm : %spending timeslip\n", pending.Hours, pending.Minutes, task)
	}

	r.printTotal(p.totalTimeWorked + r.PendingTimeslip.TotalTimeWorked())
}

// Displays the total time worked for a report
func (r Report) printTotal(totalTimeWorked int) {
	fmt.Println("===========")

	w := worked.WorkTime{}
	w.FromSeconds(totalTimeWorked)
	fmt.Printf("%4dh %3dm\n", w.Hours, w.Minutes)
}

// Prints all errors to the terminal.
func (r Report) printErrors() {
	fmt.Printf("Errors found %d:\n", len(r.errors))

	// Print the report errors
	for _, e := range r.errors {
		fmt.Printf("  %s\n", e)
	}

	// Print the Project/Task errors
	for _, p := range r.projects {
		if len(p.scanErrors) == 0 {
			continue
		}

		fmt.Printf("Project: %s:\n", p.name)
		for _, e := range p.scanErrors {
			r.errors = append(r.errors, fmt.Errorf("> %s", e.scanner))
			r.errors = append(r.errors, fmt.Errorf("  %s", e.timeslip))
		}
	}
}

func (r *Report) sortProjectsByName() {
	sort.Slice(r.projects, func(i, j int) bool {
		return strings.ToLower(r.projects[i].name) < strings.ToLower(r.projects[j].name)
	})
}

func (r Report) formattedDates() string {
	from := r.timePeriod.From().Format("Jan 2, 2006")
	to := r.timePeriod.To().Format("Jan 2, 2006")

	if from == to {
		return fmt.Sprintf("%s", from)
	} else {
		return fmt.Sprintf("%s to %s", from, to)
	}
}
