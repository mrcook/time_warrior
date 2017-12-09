// Copyright (c) 2017 Michael R. Cook

package main

import (
	"fmt"
	"os"
	"path"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/mrcook/time_warrior/configuration"
)

func init() {
	config := configuration.New()

	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	config.HomeDirectory = home

	if err := setupNewInstall(config); err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	if verifyDataFilesPresent(config) {
		fmt.Println("One or more data files are missing! Re-run the app.")
		os.Exit(1)
	}
}

func main() {
}

func setupNewInstall(config *configuration.Config) error {
	dirPath := path.Join(config.HomeDirectory, config.DataDirectory)
	if err := os.Mkdir(dirPath, 0755); os.IsNotExist(err) {
		return err
	}

	pending := path.Join(dirPath, config.Pending)
	if _, err := os.Stat(pending); err != nil {
		f, err := os.Create(pending)
		if err != nil {
			return err
		}
		defer f.Close()
		fmt.Printf("%s created!\n", config.Pending)
	}

	completed := path.Join(dirPath, config.Completed)
	if _, err := os.Stat(completed); err != nil {
		f, err := os.Create(completed)
		if err != nil {
			return err
		}
		defer f.Close()
		fmt.Printf("%s created!\n", config.Completed)
	}

	return nil
}

func verifyDataFilesPresent(config *configuration.Config) bool {
	dataPath := path.Join(config.HomeDirectory, config.DataDirectory)

	if _, err := os.Stat(dataPath); err == nil {
		return false
	}
	if _, err := os.Stat(path.Join(dataPath, config.Pending)); err == nil {
		return false
	}
	if _, err := os.Stat(path.Join(dataPath, config.Completed)); err == nil {
		return false
	}

	return true
}
