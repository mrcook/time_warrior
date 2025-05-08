// Package configuration provides some default directory and file names.
package configuration

import (
	"fmt"
	"os"
	"path"

	"github.com/mitchellh/go-homedir"
)

// Config for the application files/folders
type Config struct {
	homeDirectory   string
	dataFolder      string
	pendingFilename string
	projectFile     string
}

// New returns a new configuration with some sane defaults
func New() *Config {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return &Config{
		homeDirectory:   home,
		dataFolder:      "time_warrior",
		pendingFilename: ".pending",
		projectFile:     ".project",
	}
}

func (c Config) DataDirectoryPath() string {
	return path.Join(c.homeDirectory, c.dataFolder)
}

func (c Config) PendingFilePath() string {
	return path.Join(c.DataDirectoryPath(), c.pendingFilename)
}

func (c Config) ProjectFilePath() string {
	return path.Join(c.DataDirectoryPath(), c.projectFile)
}

func (c Config) VerifyDataFilesPresent() bool {
	if _, err := os.Stat(c.DataDirectoryPath()); err != nil {
		return false
	}

	if _, err := os.Stat(c.PendingFilePath()); err != nil {
		return false
	}

	if _, err := os.Stat(c.ProjectFilePath()); err != nil {
		return false
	}

	return true
}

// GetCurrentProject returns the current project or empty string if not set
func (c Config) GetCurrentProject() (string, error) {
	if _, err := os.Stat(c.ProjectFilePath()); err != nil {
		return "", nil
	}

	data, err := os.ReadFile(c.ProjectFilePath())
	if err != nil {
		return "", err
	}

	return string(data), nil
}

// SetCurrentProject sets the current project
func (c Config) SetCurrentProject(project string) error {
	return os.WriteFile(c.ProjectFilePath(), []byte(project), 0644)
}
