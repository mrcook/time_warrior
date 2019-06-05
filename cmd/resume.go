package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/mrcook/time_warrior/manager"
	"github.com/mrcook/time_warrior/timeslip"
)

var resumeCmd = &cobra.Command{
	Use:                   "resume",
	Short:                 "Resume a paused timeslip",
	Aliases:               []string{"r"},
	Args:                  cobra.NoArgs,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		slip, err := resumeTimeSlip()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(slip)
		}
	},
}

func init() {
	rootCmd.AddCommand(resumeCmd)
}

func resumeTimeSlip() (*timeslip.Slip, error) {
	m := manager.NewFromConfig(initializeConfig())

	slipJSON, slipError := m.PendingTimeSlip()
	if slipError != nil {
		return nil, slipError
	}

	slip := &timeslip.Slip{}
	if err := timeslip.Unmarshal(slipJSON, slip); err != nil {
		return nil, err
	}

	if err := slip.Resume(); err != nil {
		return nil, err
	}

	if err := m.SavePending(slip.ToJson()); err != nil {
		return nil, err
	}

	return slip, nil
}
