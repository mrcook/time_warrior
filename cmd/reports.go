package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/mrcook/time_warrior/manager"
	"github.com/mrcook/time_warrior/reports"
	"github.com/mrcook/time_warrior/timeslip"
)

var period = "t"

var reportCmd = &cobra.Command{
	Use:   "report [flags] PROJECT",
	Short: "Generate a report card for projects",
	Long: `Generate a report card for your projects.

Time Periods:
  - t=today, w=this week, m=this month, y=this year.
  - 1d=yesterday, 1w=last week, 1m=last month, 1y=last year.

Project name omitted: report is generated showing the total time worked for
each Project.

Project name specified: report is generated showing the total time worked
for that Project, followed by a break down of time worked for each Task.

Time Unit incorrect or missing: report is generated using *all* timeslips.

Examples:

$ tw report -p m
=> Report showing the total time worked per project for the current month.

$ tw report -p 1d MyProject
=> Report for all tasks in MyProject, with the total time worked per task,
   for yesterday.

Further instructions and examples can be found in the README.
`,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		projectName := ""
		if len(args) > 0 {
			projectName = args[0]
		}
		generateReport(projectName, period)
	},
}

func init() {
	reportCmd.Flags().StringVarP(&period, "period", "p", "", `report for the time period: t, 1d, w, m, y.`)

	rootCmd.AddCommand(reportCmd)
}

func generateReport(projectName, period string) {
	m := manager.NewFromConfig(initializeConfig())

	pendingSlip := &timeslip.Slip{}
	if pending, err := m.PendingTimeSlip(); err == nil {
		_ = timeslip.Unmarshal(pending, pendingSlip)
	}

	report := reports.New(period)
	if pendingSlip.TotalTimeWorked() > 0 {
		report.PendingTimeslip = pendingSlip
	}

	if projectName == "" {
		for _, filename := range m.AllProjectFilenames() {
			report.ProcessProjectFile(filename)
		}
	} else {
		filename, ok := m.ProjectFilename(projectName)
		if !ok {
			fmt.Println("project file not found")
			return
		}
		report.ProcessProjectFile(filename)
	}

	report.PrintReport()
}
