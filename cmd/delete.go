// Copyright (c) 2017 Michael R. Cook

package cmd

import (
	"fmt"
	"os"

	"github.com/mrcook/time_warrior/manager"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an in progress timeslip",
	Args:  cobra.NoArgs,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		if err := deletePendingTimeSlip(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Deleted!")
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

func deletePendingTimeSlip() error {
	m := manager.NewFromConfig(initializeConfig())

	if !m.PendingTimeSlipExists() {
		return fmt.Errorf("no pending timeslip found")
	}

	if err := os.Truncate(m.PendingSlipFilename(), 0); err != nil {
		return fmt.Errorf("unable to delete pending timeslip")
	}

	return nil
}
