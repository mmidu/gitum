package commands

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"

	"../app/collections"
	"../app/models"
	"../utils"
)

var setInvocation = []string{"s", "set"}

var setDescription = "Sets the user using the identifier indicated in your git-credentials.json file."

// SetCommand is the help command
var SetCommand = models.NewCommand(ExecSet, setDescription, setInvocation)

// ExecSet sets the new user
func ExecSet() {

	var identifier string
	i := 1
	if i+1 < len(os.Args) {
		identifier = os.Args[i+1]
	}

	credentialsPath := fmt.Sprintf("%s/git-credentials.json", utils.GetHomeDir())

	if utils.FileExists(credentialsPath) {

		var currentUser models.User

		filePath := fmt.Sprintf("%s/.git-credentials", utils.GetHomeDir())

		jsonFile, err := os.Open(fmt.Sprintf("%s/git-credentials.json", utils.GetHomeDir()))

		utils.Check(err)

		defer jsonFile.Close()

		var users collections.Users

		byteValue, _ := ioutil.ReadAll(jsonFile)

		json.Unmarshal(byteValue, &users)

		if identifier != "" {
			exists := users.Contains(identifier)
			if exists {
				currentUser = users.Get(identifier)
			} else {
				fmt.Println(fmt.Sprintf("%s: credentials do not exist", identifier))
				os.Exit(0)
			}
		} else {
			fmt.Println("User identifier not set.\nDefine it with the s (set) command.")
			os.Exit(0)
		}

		credentials := fmt.Sprintf("https://%s:%s@%s\n", url.QueryEscape(currentUser.Credentials.Username), url.QueryEscape(currentUser.Credentials.Password), currentUser.Credentials.Domain)

		file, err := os.Create(filePath)
		utils.Check(err)
		defer file.Close()

		w := bufio.NewWriter(file)
		_, err = w.WriteString(credentials)
		utils.Check(err)
		w.Flush()
		os.Exit(0)
	} else {
		fmt.Println("git-credentials.json file does not exists.\nGenerate it with the i (init) command.")
		os.Exit(0)
	}
}
