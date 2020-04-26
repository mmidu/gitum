package utils

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"strings"
)

// GetCurrentUser gets the currently active git user
func GetCurrentUser() map[string]string {
	credentialsFile, err := os.Open(fmt.Sprintf("%s/.git-credentials", GetHomeDir()))
	Check(err)
	defer credentialsFile.Close()

	reader := bufio.NewReader(credentialsFile)
	data, err := reader.ReadString('\n')
	Check(err)

	fl := data[8:]

	resA := strings.SplitN(fl, ":", -1)

	username := resA[0]

	resB := strings.SplitN(resA[1], "@", -1)

	password, _ := url.QueryUnescape(resB[0])

	domain := strings.TrimSuffix(resB[1], "\n")

	currentUser := map[string]string{
		"username": username,
		"password": password,
		"domain":   domain,
	}
	return currentUser
}
