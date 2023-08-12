// Package manager provides timeslip file management.
package manager

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/mrcook/time_warrior/configuration"
)

type Manager struct {
	dataDirectory string
	pendingFile   string
}

// NewFromConfig returns a new manager from a config.
func NewFromConfig(cfg *configuration.Config) *Manager {
	return &Manager{
		dataDirectory: cfg.DataDirectoryPath(),
		pendingFile:   cfg.PendingFilePath(),
	}
}

// PendingTimeSlip reads a timeslip from the pending file.
func (m Manager) PendingTimeSlip() ([]byte, error) {
	if !m.PendingTimeSlipExists() {
		return nil, fmt.Errorf("can not resume, no pending timeslip found")
	}

	slip, err := os.ReadFile(m.pendingFile)
	if err != nil {
		return nil, err
	}
	return slip, nil
}

// PendingTimeSlipExists returns true if the pending file contains a current timeslip.
func (m Manager) PendingTimeSlipExists() bool {
	file, err := os.Stat(m.pendingFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if file.Size() > 0 {
		return true
	}
	return false
}

// SaveCompleted saves a timeslip to the project JSON file.
func (m Manager) SaveCompleted(project string, slip []byte) error {
	slip = append(slip[:], []byte("\n")...)

	filename := path.Join(m.dataDirectory, toSnakeCase(project)+".json")

	file, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(slip)
	if err != nil {
		return fmt.Errorf("unable to save completed timeslip: %v", err)
	}

	return nil
}

// SavePending saves a timeslip to the pending file.
func (m Manager) SavePending(slip []byte) error {
	if len(slip) == 0 {
		return fmt.Errorf("missing pending JSON data")
	}

	err := os.WriteFile(m.pendingFile, slip, 0644)
	if err != nil {
		return fmt.Errorf("unable to save pending timeslip: %v", err)
	}

	return nil
}

// DeletePending deletes any timeslip found in the pending file.
func (m Manager) DeletePending() error {
	if err := os.Truncate(m.pendingFile, 0); err != nil {
		return fmt.Errorf("pending timeslip may not have been deleted")
	}
	return nil
}

// AllProjectFilenames return a list of file names for all projects.
func (m Manager) AllProjectFilenames() []string {
	files, _ := filepath.Glob(filepath.Join(m.dataDirectory, "*.json"))
	return files
}

// ProjectFilename returns the file name for the requested project.
func (m Manager) ProjectFilename(projectName string) (string, bool) {
	filename := filepath.Join(m.dataDirectory, toSnakeCase(projectName)+".json")

	_, err := os.Stat(filename)
	if err != nil {
		return "", false
	}

	return filename, true
}

var exp = regexp.MustCompile("([a-z0-9]+)([A-Z])")

func toSnakeCase(camel string) string {
	camel = exp.ReplaceAllString(camel, "${1}_${2}")

	return strings.ToLower(camel)
}
