package internal

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// PathTo file resolves relative or absolute paths
func PathTo(file string) string {
	if filepath.IsAbs(file) {
		return file
	}
	pwd, _ := os.Getwd()
	return path.Join(pwd, file)
}

// WriteFile and return as a string
func WriteFile(data, file string) error {
	return ioutil.WriteFile(PathTo(file), []byte(data), 0700)
}

// EnvMap is a converts the env to a map fo string[string]
func EnvMap() map[string]interface{} {
	env := make(map[string]interface{})

	for _, v := range os.Environ() {
		s := strings.Split(v, "=")
		env[s[0]] = s[1]
	}

	return env
}
