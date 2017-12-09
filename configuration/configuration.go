package configuration

// Config for the application
type Config struct {
	HomeDirectory  string
	ConfigFilename string
	DataDirectory  string
	Pending        string
	Completed      string
}

// New returns a new configuration with some sane defaults
func New() *Config {
	return &Config{
		ConfigFilename: ".tw.yml",
		DataDirectory:  "time_warrior",
		Pending:        "pending.json",
		Completed:      "completed.json",
	}
}
