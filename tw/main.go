// TimeWarrior Copyright (c) 2017-2019 Michael R. Cook
package main

import (
	"fmt"
	"os"

	"github.com/mrcook/time_warrior/cmd"
	"github.com/mrcook/time_warrior/configuration"
)

func main() {
	config := configuration.New()
	if err := setupNewInstall(config); err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
	cmd.Execute()
}

func setupNewInstall(config *configuration.Config) error {
	if config.VerifyDataFilesPresent() {
		return nil
	}

	dataFolder := config.DataDirectoryPath()
	if err := os.Mkdir(dataFolder, 0755); err == nil {
		fmt.Printf("data folder was created at %s\n", dataFolder)
	} else if !os.IsExist(err) {
		return err
	}

	// Create pending file if it doesn't exist
	pending := config.PendingFilePath()
	if _, err := os.Stat(pending); err != nil {
		f, createErr := os.Create(pending)
		if createErr != nil {
			return createErr
		}
		defer f.Close()
		fmt.Println("pending file was created!")
	}

	// Create project file if it doesn't exist
	project := config.ProjectFilePath()
	if _, err := os.Stat(project); err != nil {
		f, createErr := os.Create(project)
		if createErr != nil {
			return createErr
		}
		defer f.Close()
		fmt.Println("project file was created!")
	}

	if !config.VerifyDataFilesPresent() {
		return fmt.Errorf("one or more data files are missing! Re-run the app")
	}

	return nil
}
