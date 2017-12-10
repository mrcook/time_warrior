// Copyright (c) 2017 Michael R. Cook

package manager

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

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

func (m Manager) StartNewSlip(name string) *timeslip.Slip {
	if m.PendingTimeSlipExists() {
		slip, err := m.PendingTimeSlip()
		if err != nil {
			fmt.Errorf("%v", err)
			return nil
		}

		fmt.Println("Aborting. Pending time slip already exists!")
		return slip
	}

	slip, err := timeslip.New(name)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	m.saveAsPending(slip.ToJson())

	return slip
}

func (m Manager) DeletePendingTimeSlip() error {
	if !m.PendingTimeSlipExists() {
		return fmt.Errorf("no pending time slip found")
	}

	if err := os.Truncate(m.pendingFile, 0); err != nil {
		return fmt.Errorf("unable to delete pending time slip")
	}

	return nil
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

func (m Manager) saveAsPending(slipJson []byte) {
	if len(slipJson) == 0 {
		fmt.Println("Missing pending JSON data. Aborting!")
		return
	}

	err := ioutil.WriteFile(m.pendingFile, slipJson, 0644)
	if err != nil {
		fmt.Errorf("unable to save pending JSON data: %v", err)
	}
}
