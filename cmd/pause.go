package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/mrcook/time_warrior/manager"
	"github.com/mrcook/time_warrior/timeslip"
)

var pauseCmd = &cobra.Command{
	Use:                   "pause",
	Short:                 "Pause a started timeslip",
	Aliases:               []string{"p"},
	Args:                  cobra.NoArgs,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		slip, err := pauseTimeSlip()

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(slip)
		}
	},
}

func init() {
	rootCmd.AddCommand(pauseCmd)
}

func pauseTimeSlip() (*timeslip.Slip, error) {
	m := manager.NewFromConfig(initializeConfig())

	slipJSON, slipError := m.PendingTimeSlip()
	if slipError != nil {
		return nil, slipError
	}

	if len(slipJSON) == 0 {
		return nil, fmt.Errorf("no timeslip to pause")
	}

	slip := &timeslip.Slip{}
	if err := timeslip.Unmarshal(slipJSON, slip); err != nil {
		return nil, err
	}

	if err := slip.Pause(); err != nil {
		return nil, err
	}

	if err := m.SavePending(slip.ToJson()); err != nil {
		return nil, err
	}

	return slip, nil
}
