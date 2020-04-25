package controllers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"../models"
	"../utils"
)

// JSONStub is the default json structure of the git-credentials.json file
const JSONStub = "{\n\t\"users\": [\n\t\t{\n\t\t\t\"identifier\": \"github-work\",\n\t\t\t\"credentials\": {\n\t\t\t\t\"username\": \"username\",\n\t\t\t\t\"password\": \"password1\",\n\t\t\t\t\"domain\": \"github.com\"\n\t\t\t}\n\t\t},\n\t\t{\n\t\t\t\"identifier\": \"two\",\n\t\t\t\"credentials\": {\n\t\t\t\t\"username\": \"uname2\",\n\t\t\t\t\"password\": \"password2\",\n\t\t\t\t\"domain\": \"domain\"\n\t\t\t}\n\t\t}\n\t]\n}"

// Init initializes the JSON stub file
func Init() {
	path := fmt.Sprintf("%s/git-credentials.json", utils.GetHomeDir())
	if !utils.FileExists(path) {
		file, err := os.Create(path)
		utils.Check(err)
		defer file.Close()

		w := bufio.NewWriter(file)
		_, err = w.WriteString(JSONStub)
		utils.Check(err)
		w.Flush()
	} else {
		fmt.Println(fmt.Sprintf("%s/git-credentials.json file already exists.", utils.GetHomeDir()))
	}

	os.Exit(0)
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
		case "-lu", "--listUsers":
			printCredentials()
		case "-c", "--current":
			getCurrent()
		case "-h", "--help":
			fmt.Println("HELP STRING")
			os.Exit(0)
		}
	}
	return identifier
}

func printCredentials() {
	jsonFile, err := os.Open(fmt.Sprintf("%s/git-credentials.json", utils.GetHomeDir()))
	utils.Check(err)
	defer jsonFile.Close()

	var users models.Users

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &users)

	fmt.Println(users.List())

	os.Exit(0)
}

func getCurrent() {
	credentialsFile, err := os.Open(fmt.Sprintf("%s/.git-credentials", utils.GetHomeDir()))
	utils.Check(err)
	defer credentialsFile.Close()

	reader := bufio.NewReader(credentialsFile)
	data, err := reader.ReadString('\n')
	utils.Check(err)
	fmt.Printf(string(data))

	os.Exit(0)
}
