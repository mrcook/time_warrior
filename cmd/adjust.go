// Copyright (c) 2017 Michael R. Cook

package cmd

import (
	"fmt"

	"github.com/mrcook/time_warrior/manager"
	"github.com/spf13/cobra"
)

var negative bool

var adjustCmd = &cobra.Command{
	Use:   "adjust DURATION",
	Short: "Adjust +/- the time worked on a timeslip",
	Long: `Increase or decrease the time worked on a paused timeslip using a
duration string based on time units of hours, minutes, or seconds.

The DURATION string should be in the format of '10m' - a decimal
number followed by a single time unit character (no spaces).

Allowed units are 'h', 'm', and 's'.

Example strings: '72m', '2h', '130s', '30m', '720s'

To subtract a value, specify the -n (negative) flag.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		m := manager.NewFromConfig(initializeConfig())

		var duration string
		if negative == true {
			duration = "-" + args[0]
		} else {
			duration = args[0]
		}

		slip, err := m.Adjust(duration)
		if slip != nil {
			fmt.Println(slip)
		}
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	adjustCmd.Flags().BoolVarP(&negative, "negative", "n", false, "use negative time duration")

	rootCmd.AddCommand(adjustCmd)
}
