// Copyright (c) 2017 Michael R. Cook

package cmd

import (
	"fmt"

	"github.com/mrcook/time_warrior/manager"
	"github.com/spf13/cobra"
)

var resumeCmd = &cobra.Command{
	Use:     "resume",
	Short:   "Resume a paused time slip",
	Aliases: []string{"r"},
	Run: func(cmd *cobra.Command, args []string) {
		m := manager.NewFromConfig(initializeConfig())
		slip, err := m.ResumeTimeSlip()
		if err != nil {
			fmt.Errorf("%v", err)
		} else {
			fmt.Println(slip)
		}
	},
}

func init() {
	rootCmd.AddCommand(resumeCmd)
}
