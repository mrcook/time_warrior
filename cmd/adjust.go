package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/mrcook/time_warrior/manager"
	"github.com/mrcook/time_warrior/timeslip"
	"github.com/mrcook/time_warrior/timeslip/status"
)

var adjustNegative bool

var adjustCmd = &cobra.Command{
	Use:   "adjust DURATION",
	Short: "Adjust the time worked on a timeslip",
	Long: `Increase or decrease the time worked on a timeslip using a
duration string based on time units of hours, minutes, or seconds.

The DURATION string should be in the format of '10m' - a decimal
number followed by a single time unit character (no spaces).

Allowed units are 'h', 'm', and 's'.

Example strings: '72m', '2h', '130s', '30m', '720s'

To subtract a value, specify the -n (negative) flag.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var duration string
		if adjustNegative == true {
			duration = "-" + args[0]
		} else {
			duration = args[0]
		}

		slip, err := adjust(duration)

		if err != nil {
			fmt.Println(err)
		}

		if slip != nil {
			fmt.Println(slip)
		}
	},
}

func init() {
	adjustCmd.Flags().BoolVarP(&adjustNegative, "negative", "n", false, "use negative time duration")

	rootCmd.AddCommand(adjustCmd)
}

// Adjust a pending timeslip +/- a given amount
// Only paused timeslips should be adjusted
func adjust(adjustment string) (*timeslip.Slip, error) {
	m := manager.NewFromConfig(initializeConfig())

	if !m.PendingTimeSlipExists() {
		return nil, fmt.Errorf("no pending timeslip found")
	}

	slipJSON, slipError := m.PendingTimeSlip()
	if slipError != nil {
		return nil, slipError
	}

	slip := &timeslip.Slip{}
	if err := timeslip.Unmarshal(slipJSON, slip); err != nil {
		return nil, err
	}

	// Running timeslips must be paused
	running := false
	if slip.Status == status.Started() {
		running = true
		slip.Pause()
	}

	if err := slip.Adjust(adjustment); err != nil {
		return nil, err
	}

	// Resume timeslip if it was previously running
	if running {
		slip.Resume()
	}

	if err := m.SavePending(slip.ToJson()); err != nil {
		return slip, fmt.Errorf("timeslip may not have been saved: %v", err)
	}

	return slip, nil
}
