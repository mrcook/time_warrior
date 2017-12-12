// Copyright (c) 2017 Michael R. Cook

package cmd

import (
	"fmt"

	"github.com/mrcook/time_warrior/manager"
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
		m := manager.NewFromConfig(initializeConfig())

		slip, err := m.Done(args[0])
		if slip != nil {
			fmt.Println(slip)
		}
		if err != nil {
			fmt.Errorf("%v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
