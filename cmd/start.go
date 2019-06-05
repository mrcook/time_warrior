package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/mrcook/time_warrior/manager"
	"github.com/mrcook/time_warrior/timeslip"
)

var startCmd = &cobra.Command{
	Use:   "start Project.Task",
	Short: "Start a new timeslip",
	Long: `Start working on a new task, providing a project, and optional task name.

Only alphanumeric characters are allowed - no spaces - the project and task
name must be separated by a period. Example: MyProject.StartTask`,
	Aliases:               []string{"s"},
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		slip, err := startNewSlip(args[0])

		if err != nil {
			fmt.Println(err)
		}

		if slip != nil {
			fmt.Println(slip)
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}

func startNewSlip(name string) (*timeslip.Slip, error) {
	m := manager.NewFromConfig(initializeConfig())

	if m.PendingTimeSlipExists() {
		slipJSON, slipError := m.PendingTimeSlip()
		if slipError == nil {
			return nil, fmt.Errorf("pending timeslip already exists")
		}

		slip := &timeslip.Slip{}
		if err := timeslip.Unmarshal(slipJSON, slip); err != nil {
			return nil, err
		}

		return slip, nil
	}

	slip, err := timeslip.New(name)
	if err != nil {
		return nil, err
	}

	if err := m.SavePending(slip.ToJson()); err != nil {
		return nil, err
	}

	return slip, nil
}
