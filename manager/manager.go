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

func (m Manager) StartNewSlip(name string) {
	if m.outstandingTimeSlipExists() {
		pending, err := m.outstandingTimeSlip()
		if err != nil {
			fmt.Errorf("%v", err)
			return
		}

		fmt.Println("Aborting. Outstanding time slip present!")
		fmt.Println(pending.String())
		return
	}

	slip, err := timeslip.New(name)
	if err != nil {
		fmt.Println(err)
		return
	}

	m.saveAsPending(slip.ToJson())
	fmt.Println(slip)
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

func (m Manager) outstandingTimeSlipExists() bool {
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

func (m Manager) outstandingTimeSlip() (*timeslip.Slip, error) {
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
