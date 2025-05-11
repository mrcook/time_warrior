package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/mrcook/time_warrior/manager"
	"github.com/mrcook/time_warrior/reports"
)

var agendaCmd = &cobra.Command{
	Use:   "agenda [flags] PROJECT",
	Short: "Generate an agenda view of your time tracking data",
	Long: `Generate a Mermaid diagram showing your time tracking data in an agenda view.
This view helps visualize how your time was spent across different tasks and projects.

Time Periods:
  - t=today, w=this week, m=this month, y=this year.
  - 1d=yesterday, 1w=last week, 1m=last month, 1y=last year.

Project name omitted: agenda is generated showing all projects.
Project name specified: agenda is generated for that specific project.

Examples:
  $ tw agenda -p m
  => Agenda view showing all projects for the current month.

  $ tw agenda -p 1d MyProject
  => Agenda view for MyProject for yesterday.`,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		projectName := ""
		if len(args) > 0 {
			projectName = args[0]
		}
		generateAgenda(projectName, period)
	},
}

func init() {
	agendaCmd.Flags().StringVarP(&period, "period", "p", "", `report for the time period: t, 1d, w, m, y.`)
	rootCmd.AddCommand(agendaCmd)
}

func generateAgenda(projectName, period string) {
	m := manager.NewFromConfig(initializeConfig())

	report := reports.New(period)

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

	// Generate Mermaid diagram
	fmt.Println("```mermaid")
	fmt.Println("gantt")
	fmt.Println("    title Time Tracking Agenda")
	fmt.Println("    dateFormat  YYYY-MM-DD HH:mm")
	fmt.Println("    axisFormat %H:%M")

	// Add tasks to the diagram
	for _, p := range report.Projects() {
		for _, t := range p.SortedTasks() {
			startTime := time.Unix(int64(t.Started()), 0)
			endTime := time.Unix(int64(t.Finished()), 0)

			// If the task is not finished, use current time as end time
			if t.Finished() == 0 {
				endTime = time.Now()
			}

			// Format the task name to be Mermaid-compatible
			taskName := fmt.Sprintf("%s.%s", p.Name(), t.Name())
			if t.Name() == "." {
				taskName = p.Name()
			}

			// Add start time to task name to distinguish between multiple instances
			if t.Started() > 0 {
				taskName = fmt.Sprintf("%s (%s)", taskName, startTime.Format("15:04"))
			}

			fmt.Printf("    %s :%s, %s, %s\n",
				taskName,
				taskName,
				startTime.Format("2006-01-02 15:04"),
				endTime.Format("2006-01-02 15:04"))
		}
	}

	fmt.Println("```")
}
