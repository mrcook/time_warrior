package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var projectCmd = &cobra.Command{
	Use:   "project [project]",
	Short: "Set or show the current project",
	Long: `Set or show the current project for time tracking.
If no project is provided, the current project will be displayed.
If a project is provided, it will be set as the current project.`,
	Aliases: []string{"pr"},
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		config := initializeConfig()

		if len(args) == 0 {
			// Show current project
			proj, err := config.GetCurrentProject()
			if err != nil {
				fmt.Println(err)
				return
			}
			if proj == "" {
				fmt.Println("No project set")
			} else {
				fmt.Printf("Current project: %s\n", proj)
			}
			return
		}

		// Set new project
		project := args[0]
		if err := config.SetCurrentProject(project); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Project set to: %s\n", project)
	},
}

func init() {
	rootCmd.AddCommand(projectCmd)
}
