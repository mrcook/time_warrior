// Package configuration provides some default directory and file names.
package configuration

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
)

// Config for the application
type Config struct {
	HomeDirectory string
	DataDirectory string
	Pending       string
}

// New returns a new configuration with some sane defaults
func New() *Config {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return &Config{
		HomeDirectory: home,
		DataDirectory: "time_warrior",
		Pending:       ".pending",
	}
}
