package config

import (
	"os"
	"path/filepath"
)

// DefaultDir returns the default directory to store configuration files
func DefaultDir() string {
	if configHome, ok := os.LookupEnv("XDG_CONFIG_HOME"); ok {
		return filepath.Join(configHome, "clipboard-sync")
	}
	return filepath.Join(os.Getenv("HOME"), ".config", "clipboard-sync")
}
