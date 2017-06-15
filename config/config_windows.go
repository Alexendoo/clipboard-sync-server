package config

import (
	"os"
	"path/filepath"
)

// DefaultDir returns the default directory to store configuration files
func DefaultDir() string {
	return filepath.Join(os.Getenv("LOCALAPPDATA"), "Clipboard Sync")
}
