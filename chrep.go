package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
}

func getHomeDir() string {
	cuser, err := user.Current()
	if err != nil {
		panic(err)
	}
	return cuser.HomeDir
}

func main() {

	filePath := fmt.Sprintf("%s/.git-credentials", getHomeDir())
	fmt.Println(filePath)
	hasArg := len(os.Args) > 1
	var identifier string

	if hasArg {
		identifier = os.Args[1]
	}

	jsonFile, err := os.Open("credentials.json")

	if err != nil {
		fmt.Println(err)
	}

	var users Users

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &users)

	if identifier != "" {
		exists := users.Contains(identifier)
		if exists {
			currentUser := users.Get(identifier)
			fmt.Println(currentUser)
		} else {
			fmt.Println(fmt.Sprintf("%s: credentials do not exist", identifier))
		}
	} else {
		currentUser := users.Users[0]
		fmt.Println(currentUser)
	}
}
