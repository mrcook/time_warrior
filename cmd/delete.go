package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/mrcook/time_warrior/manager"
)

var deleteCmd = &cobra.Command{
	Use:                   "delete",
	Short:                 "Delete an in progress timeslip",
	Args:                  cobra.NoArgs,
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

	err := m.DeletePending()
	return err
}
