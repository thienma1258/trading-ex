package config

import (
	"path/filepath"
)

type (
	Config struct {
		App        App    `mapstructure:"app"`
		TimeFormat string `mapstructure:"time_format"`
	}

	App struct {
		LogLevel string `mapstructure:"log_level"`
	}
)

// Default is a configuration instance.
var Default = Config{
	App: App{
		LogLevel: "debug",
	},
} // nolint:gochecknoglobals // config must be global

// SetConfig reads a config file and returs an initialized config instance.
func SetConfig(confPath string) error {
	confPath, err := filepath.Abs(confPath)
	if err != nil {
		return err
	}

	return nil
}
