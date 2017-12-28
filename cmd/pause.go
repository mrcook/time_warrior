// Copyright (c) 2017 Michael R. Cook

package cmd

import (
	"fmt"

	"github.com/mrcook/time_warrior/manager"
	"github.com/mrcook/time_warrior/timeslip"
	"github.com/spf13/cobra"
)

var pauseCmd = &cobra.Command{
	Use:     "pause",
	Short:   "Pause a started timeslip",
	Aliases: []string{"p"},
	Args:    cobra.NoArgs,
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

	slip, err := m.PendingTimeSlip()
	if err != nil {
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
