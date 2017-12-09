// Copyright (c) 2017 Michael R. Cook

package cmd

import (
	"fmt"

	"github.com/mrcook/time_warrior/manager"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start PROJECT.TASK",
	Short: "Start a new time slip",
	Long: `Start working on a new task, providing a project, and optional task name.

Only alphanumeric characters are allowed - no spaces - the project and task
name must be separated by a period. Example: TimeWarrior.StartTask`,
	Aliases: []string{"s"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		m := manager.NewFromConfig(initializeConfig())
		if slip := m.StartNewSlip(args[0]); slip != nil {
			fmt.Println(slip)
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
