package main

import (
	"os"
	"strings"
)

// ExpandPath expands home directory path for root key in TOML config
func ExpandPath(path string) string {
	if strings.HasPrefix(path, "~/") {
		home, err := os.UserHomeDir()
		if err != nil {
			return path
		}
		return strings.Replace(path, "~", home, 1)
	}
	return path
}
