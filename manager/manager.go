// Copyright (c) 2017 Michael R. Cook

package manager

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/mrcook/time_warrior/configuration"
	"github.com/mrcook/time_warrior/timeslip"
)

type Manager struct {
	dataDirectory string
	pendingFile   string
}

func NewFromConfig(cfg *configuration.Config) *Manager {
	return &Manager{
		dataDirectory: path.Join(cfg.HomeDirectory, cfg.DataDirectory),
		pendingFile:   path.Join(cfg.HomeDirectory, cfg.DataDirectory, cfg.Pending),
	}
}

func (m Manager) PendingSlipFilename() string {
	return m.pendingFile
}

func (m Manager) StartNewSlip(name string) *timeslip.Slip {
	if m.PendingTimeSlipExists() {
		slip, err := m.PendingTimeSlip()
		if err != nil {
			fmt.Errorf("%v", err)
			return nil
		}

		fmt.Println("Aborting. Pending timeslip already exists!")
		return slip
	}

	slip, err := timeslip.New(name)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	if err := m.saveAsPending(slip.ToJson()); err != nil {
		fmt.Errorf("%v", err)
		return nil
	}

	return slip
}

func (m Manager) PauseTimeSlip() (*timeslip.Slip, error) {
	slip, err := m.PendingTimeSlip()
	if err != nil {
		return nil, err
	}

	if err := slip.Pause(); err != nil {
		return nil, err
	}

	if err := m.saveAsPending(slip.ToJson()); err != nil {
		return nil, err
	}

	return slip, nil
}

func (m Manager) ResumeTimeSlip() (*timeslip.Slip, error) {
	slip, err := m.PendingTimeSlip()
	if err != nil {
		return nil, err
	}

	if err := slip.Resume(); err != nil {
		return nil, err
	}

	if err := m.saveAsPending(slip.ToJson()); err != nil {
		return nil, err
	}

	return slip, nil
}

func (m Manager) DeletePendingTimeSlip() error {
	if !m.PendingTimeSlipExists() {
		return fmt.Errorf("no pending timeslip found")
	}

	if err := os.Truncate(m.pendingFile, 0); err != nil {
		return fmt.Errorf("unable to delete pending timeslip")
	}

	return nil
}

func (m Manager) Done(description string) (*timeslip.Slip, error) {
	if !m.PendingTimeSlipExists() {
		return nil, fmt.Errorf("no pending timeslip found")
	}

	slip, err := m.PendingTimeSlip()
	if err != nil {
		return nil, err
	}

	slip.Done(description)

	if err := m.saveCompletedSlip(slip); err != nil {
		return slip, err
	}

	if err := os.Truncate(m.PendingSlipFilename(), 0); err != nil {
		return slip, fmt.Errorf("pending timeslip may not have been deleted")
	}

	return slip, nil
}

func (m Manager) PendingTimeSlip() (*timeslip.Slip, error) {
	record, err := ioutil.ReadFile(m.pendingFile)
	if err != nil {
		return nil, err
	}

	slip, err := timeslip.NewFromJSON(record)
	if err != nil {
		return nil, err
	}

	return slip, nil
}

func (m Manager) PendingTimeSlipExists() bool {
	file, err := os.Stat(m.pendingFile)
	if err != nil {
		fmt.Errorf("%v", err)
		os.Exit(1)
	}

	if file.Size() > 0 {
		return true
	}
	return false
}

func (m Manager) saveAsPending(slipJson []byte) error {
	if len(slipJson) == 0 {
		return fmt.Errorf("missing pending JSON data")
	}

	err := ioutil.WriteFile(m.pendingFile, slipJson, 0644)
	if err != nil {
		return fmt.Errorf("unable to save pending JSON data: %v", err)
	}

	return nil
}

func (m Manager) saveCompletedSlip(slip *timeslip.Slip) error {
	slipJson := slip.ToJson()
	slipJson = append(slipJson[:], []byte("\n")...)

	project := toSnakeCase(slip.Project) + ".json"
	filename := path.Join(m.dataDirectory, project)

	file, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(slipJson)
	if err != nil {
		return fmt.Errorf("unable to save pending JSON data: %v", err)
	}

	return nil
}

var exp = regexp.MustCompile("([a-z0-9])([A-Z0-9])")

func toSnakeCase(camel string) string {
	camel = exp.ReplaceAllString(camel, "${1}_${2}")

	return strings.ToLower(camel)
}
