// Copyright (c) 2017 Michael R. Cook

package cmd

import (
	"fmt"

	"github.com/mrcook/time_warrior/manager"
	"github.com/spf13/cobra"
)

var pauseCmd = &cobra.Command{
	Use:     "pause",
	Short:   "Pause a started timeslip",
	Aliases: []string{"p"},
	Args:    cobra.NoArgs,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		m := manager.NewFromConfig(initializeConfig())
		slip, err := m.PauseTimeSlip()
		if err != nil {
			fmt.Errorf("%v", err)
		} else {
			fmt.Println(slip)
		}
	},
}

func init() {
	rootCmd.AddCommand(pauseCmd)
}
