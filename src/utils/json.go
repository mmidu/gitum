package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"../app/collections"
)

// GetCredentials gets the currently saved credentials
func GetCredentials() collections.Users {
	jsonFile, err := os.Open(fmt.Sprintf("%s/git-credentials.json", GetHomeDir()))
	Check(err)
	defer jsonFile.Close()

	var users collections.Users

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &users)

	return users
}
