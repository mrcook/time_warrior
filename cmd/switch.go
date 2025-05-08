package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/mrcook/time_warrior/manager"
	"github.com/mrcook/time_warrior/timeslip"
)

var switchCmd = &cobra.Command{
	Use:   "switch [Project.Task]",
	Short: "Switch from current task to a new task",
	Long: `Stop the current task and start a new one in a single command.
If no project is provided in the task name, the current project will be used.
If no project is set, you must provide the full project name.`,
	Aliases:               []string{"sw"},
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		slip, err := switchTask(args[0])

		if err != nil {
			fmt.Println(err)
		}

		if slip != nil {
			fmt.Println(slip)
		}
	},
}

func init() {
	rootCmd.AddCommand(switchCmd)
}

func switchTask(name string) (*timeslip.Slip, error) {
	m := manager.NewFromConfig(initializeConfig())

	// First, check if there's a current task to stop
	if m.PendingTimeSlipExists() {
		slipJSON, err := m.PendingTimeSlip()
		if err != nil {
			return nil, err
		}

		slip := &timeslip.Slip{}
		if err := timeslip.Unmarshal(slipJSON, slip); err != nil {
			return nil, err
		}

		// Mark the current task as done
		slip.Done("Switched to new task")
		if err := m.SaveCompleted(slip.Project, slip.ToJson()); err != nil {
			return nil, err
		}

		// Delete the pending timeslip
		if err := m.DeletePending(); err != nil {
			return nil, err
		}
	}

	// Now start the new task
	return startNewSlip(name)
}
