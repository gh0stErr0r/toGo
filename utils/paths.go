package utils

import (
	"os"
	"path/filepath"
	"strings"
)

func ExpandHome(path string) string {
	if strings.HasPrefix(path, "~/") {
		homeDir, err := os.UserHomeDir()
		if err == nil {
			return filepath.Join(homeDir, path[2:])
		}
	}
	return path
}
