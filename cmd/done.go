package cmd

import (
	"fmt"

	"github.com/mrcook/time_warrior/manager"
	"github.com/mrcook/time_warrior/timeslip"
	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:                   "done 'Description'",
	Short:                 "Mark current timeslip as completed",
	Long:                  `Mark the current timeslip as done, providing a useful description.`,
	Aliases:               []string{"d"},
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		slip, err := done(args[0])

		if err != nil {
			fmt.Println(err)
		}

		if slip != nil {
			fmt.Println(slip)
		}
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}

func done(description string) (*timeslip.Slip, error) {
	m := manager.NewFromConfig(initializeConfig())

	if !m.PendingTimeSlipExists() {
		return nil, fmt.Errorf("no pending timeslip found")
	}

	slipJSON, slipError := m.PendingTimeSlip()
	if slipError != nil {
		return nil, slipError
	}

	slip := &timeslip.Slip{}
	if err := timeslip.Unmarshal(slipJSON, slip); err != nil {
		return nil, err
	}

	slip.Done(description)

	if err := m.SaveCompleted(slip.Project, slip.ToJson()); err != nil {
		return slip, err
	}

	if err := m.DeletePending(); err != nil {
		return slip, err
	}

	return slip, nil
}
