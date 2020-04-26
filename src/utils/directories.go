package utils

import "os/user"

// GetHomeDir gets the home directory
func GetHomeDir() string {
	cuser, err := user.Current()
	Check(err)
	return cuser.HomeDir
}
