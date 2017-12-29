// Package manager provides timeslip file management.
package manager

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/mrcook/time_warrior/configuration"
)

type manager struct {
	dataDirectory string
	pendingFile   string
}

// NewFromConfig returns a new manager from a config.
func NewFromConfig(cfg *configuration.Config) *manager {
	return &manager{
		dataDirectory: path.Join(cfg.HomeDirectory, cfg.DataDirectory),
		pendingFile:   path.Join(cfg.HomeDirectory, cfg.DataDirectory, cfg.Pending),
	}
}

// PendingTimeSlip reads a timeslip from the pending file.
func (m manager) PendingTimeSlip() ([]byte, error) {
	slip, err := ioutil.ReadFile(m.pendingFile)
	if err != nil {
		return nil, err
	}
	return slip, nil
}

// PendingTimeSlipExists returns true if the pending file contains a current timeslip.
func (m manager) PendingTimeSlipExists() bool {
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
func (m manager) SaveCompleted(project string, slip []byte) error {
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
func (m manager) SavePending(slip []byte) error {
	if len(slip) == 0 {
		return fmt.Errorf("missing pending JSON data")
	}

	err := ioutil.WriteFile(m.pendingFile, slip, 0644)
	if err != nil {
		return fmt.Errorf("unable to save pending timeslip: %v", err)
	}

	return nil
}

// DeletePending deletes any timeslip found in the pending file.
func (m manager) DeletePending() error {
	if err := os.Truncate(m.pendingFile, 0); err != nil {
		return fmt.Errorf("pending timeslip may not have been deleted")
	}
	return nil
}

var exp = regexp.MustCompile("([a-z0-9])([A-Z0-9])")

func toSnakeCase(camel string) string {
	camel = exp.ReplaceAllString(camel, "${1}_${2}")

	return strings.ToLower(camel)
}
