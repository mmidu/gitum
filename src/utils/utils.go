package utils

import (
	"os"
	"os/user"
)

// Check checks errors
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// GetHomeDir gets the home directory
func GetHomeDir() string {
	cuser, err := user.Current()
	Check(err)
	return cuser.HomeDir
}

// FileExists checks whether a file exists
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
