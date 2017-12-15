// Copyright (c) 2017 Michael R. Cook

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/mrcook/time_warrior/configuration"
	"github.com/mrcook/time_warrior/manager"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "tw",
	Version: "0.0.1",
	Short:   "TimeWarrior: a CLI based time tracking tool",
	Long: `TimeWarrior is a command line time tracking tool for developers and freelance
workers who need to track time worked on their client and personal projects.`,
	Run: func(cmd *cobra.Command, args []string) {
		m := manager.NewFromConfig(initializeConfig())
		if !m.PendingTimeSlipExists() {
			cmd.Help()
			return
		}

		slip, err := m.PendingTimeSlip()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(slip)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// initializeConfig reads in config file and ENV variables if set.
func initializeConfig() *configuration.Config {
	return configuration.New()
}
