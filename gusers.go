package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"os/user"
)

// Users is users list struct
type Users struct {
	Users []User `json:"users"`
}

// Contains checks whether the Users list contains the passed identifier
func (u Users) Contains(identifier string) bool {
	for _, user := range u.Users {
		if user.Identifier == identifier {
			return true
		}
	}
	return false
}

// Get retrieves a user by its identifier
func (u Users) Get(identifier string) User {
	var user User
	for _, user := range u.Users {
		if user.Identifier == identifier {
			return user
		}
	}
	return user
}

// User is the strict that holds the user data
type User struct {
	Identifier  string      `json:"identifier"`
	Credentials Credentials `json:"credentials"`
}

// Credentials is the struct that holds the user credentials
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Domain   string `json:"domain"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getHomeDir() string {
	cuser, err := user.Current()
	check(err)
	return cuser.HomeDir
}

func initialize() {
	file, err := os.Create(fmt.Sprintf("%s/git-credentials.json", getHomeDir()))
	check(err)
	defer file.Close()

	w := bufio.NewWriter(file)
	_, err = w.WriteString(JSONStub)
	check(err)
	w.Flush()
}

func manageArguments() string {
	var identifier string
	for i := 0; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "-i", "--init":
			initialize()
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

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// JSONStub is the default json structure of the git-credentials.json file
const JSONStub = "{\n\t\"users\": [\n\t\t{\n\t\t\t\"identifier\": \"github-work\",\n\t\t\t\"credentials\": {\n\t\t\t\t\"username\": \"username\",\n\t\t\t\t\"password\": \"password1\",\n\t\t\t\t\"domain\": \"github.com\"\n\t\t\t}\n\t\t},\n\t\t{\n\t\t\t\"identifier\": \"two\",\n\t\t\t\"credentials\": {\n\t\t\t\t\"username\": \"uname2\",\n\t\t\t\t\"password\": \"password2\",\n\t\t\t\t\"domain\": \"domain\"\n\t\t\t}\n\t\t}\n\t]\n}"

func main() {

	var identifier string

	identifier = manageArguments()

	credentialsPath := fmt.Sprintf("%s/git-credentials.json", getHomeDir())

	if fileExists(credentialsPath) {

		var currentUser User

		filePath := fmt.Sprintf("%s/.git-credentials", getHomeDir())

		jsonFile, err := os.Open(fmt.Sprintf("%s/git-credentials.json", getHomeDir()))

		check(err)

		var users Users

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
		check(err)
		defer file.Close()

		w := bufio.NewWriter(file)
		_, err = w.WriteString(credentials)
		check(err)
		w.Flush()
		fmt.Println(credentials)
	} else {
		fmt.Println("git-credentials.json file does not exists.\nGenerate it with the -i (--init) flag.")
		return
	}
}
