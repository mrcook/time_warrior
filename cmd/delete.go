// Copyright (c) 2017 Michael R. Cook

package cmd

import (
	"fmt"

	"github.com/mrcook/time_warrior/manager"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an in progress timeslip",
	Args:  cobra.NoArgs,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		m := manager.NewFromConfig(initializeConfig())
		if err := m.DeletePendingTimeSlip(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Deleted!")
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
