// Copyright (c) 2017 Michael R. Cook

package cmd

import (
	"fmt"
	"os"

	"github.com/mrcook/time_warrior/manager"
	"github.com/mrcook/time_warrior/timeslip"
	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:     "done 'Description'",
	Short:   "Mark current timeslip as completed",
	Long:    `Mark the current timeslip as done, providing a useful description.`,
	Aliases: []string{"d"},
	Args:    cobra.ExactArgs(1),
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

	slip, err := m.PendingTimeSlip()
	if err != nil {
		return nil, err
	}

	slip.Done(description)

	if err := m.SaveCompletedSlip(slip); err != nil {
		return slip, err
	}

	if err := os.Truncate(m.PendingSlipFilename(), 0); err != nil {
		return slip, fmt.Errorf("pending timeslip may not have been deleted")
	}

	return slip, nil
}
