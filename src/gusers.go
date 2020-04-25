package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"

	"./controllers"
	"./models"
	"./utils"
)

func main() {

	var identifier string

	identifier = controllers.ManageArguments()

	credentialsPath := fmt.Sprintf("%s/git-credentials.json", utils.GetHomeDir())

	if utils.FileExists(credentialsPath) {

		var currentUser models.User

		filePath := fmt.Sprintf("%s/.git-credentials", utils.GetHomeDir())

		jsonFile, err := os.Open(fmt.Sprintf("%s/git-credentials.json", utils.GetHomeDir()))

		utils.Check(err)

		var users models.Users

		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)

		json.Unmarshal(byteValue, &users)

		if identifier != "" {
			exists := users.Contains(identifier)
			if exists {
				currentUser = users.Get(identifier)
			} else {
				fmt.Println(fmt.Sprintf("%s: credentials do not exist", identifier))
				return
			}
		} else {
			fmt.Println("User identifier not set.\nDefine it with the -s (--set) flag.")
			return
		}

		credentials := fmt.Sprintf("https://%s:%s@%s\n", url.QueryEscape(currentUser.Credentials.Username), url.QueryEscape(currentUser.Credentials.Password), currentUser.Credentials.Domain)

		file, err := os.Create(filePath)
		utils.Check(err)
		defer file.Close()

		w := bufio.NewWriter(file)
		_, err = w.WriteString(credentials)
		utils.Check(err)
		w.Flush()
		fmt.Println(credentials)
	} else {
		fmt.Println("git-credentials.json file does not exists.\nGenerate it with the -i (--init) flag.")
		return
	}
}
