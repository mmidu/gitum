package controllers

import (
	"bufio"
	"fmt"
	"os"

	"../utils"
)

// JSONStub is the default json structure of the git-credentials.json file
const JSONStub = "{\n\t\"users\": [\n\t\t{\n\t\t\t\"identifier\": \"github-work\",\n\t\t\t\"credentials\": {\n\t\t\t\t\"username\": \"username\",\n\t\t\t\t\"password\": \"password1\",\n\t\t\t\t\"domain\": \"github.com\"\n\t\t\t}\n\t\t},\n\t\t{\n\t\t\t\"identifier\": \"two\",\n\t\t\t\"credentials\": {\n\t\t\t\t\"username\": \"uname2\",\n\t\t\t\t\"password\": \"password2\",\n\t\t\t\t\"domain\": \"domain\"\n\t\t\t}\n\t\t}\n\t]\n}"

// Init initializes the JSON stub file
func Init() {
	file, err := os.Create(fmt.Sprintf("%s/git-credentials.json", utils.GetHomeDir()))
	utils.Check(err)
	defer file.Close()

	w := bufio.NewWriter(file)
	_, err = w.WriteString(JSONStub)
	utils.Check(err)
	w.Flush()
}

// ManageArguments get arguments from the command line and executes the right functions
func ManageArguments() string {
	var identifier string
	for i := 0; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "-i", "--init":
			Init()
		case "-s", "--set":
			if i+1 < len(os.Args) {
				identifier = os.Args[i+1]
			}
		case "-h", "--help":
			fmt.Println("HELP STRING")
		}
	}
	return identifier
}
