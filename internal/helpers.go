package internal

import (
	"os"
	"path"
	"path/filepath"
)

// PathTo file resolves relative or absolute paths
func PathTo(file string) string {
	if filepath.IsAbs(file) {
		return file
	}
	pwd, _ := os.Getwd()
	return path.Join(pwd, file)
}
