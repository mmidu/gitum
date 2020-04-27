package commands

import (
	"encoding/json"
	"fmt"

	"../app/models"
	"../utils"
)

var addInvocation = []string{"a", "add"}

var addDescription = "Adds a user."

// AddCommand is the add command
var AddCommand = models.NewCommand(ExecAdd, addDescription, addInvocation)

// ExecAdd executed the add command
func ExecAdd(){

	identifier := utils.GetData("identifier")

	username := utils.GetData("username")

	password := utils.GetData("password")

	domain := utils.GetData("domain")

	credentials := models.NewCredentials(username, password, domain)

	user := models.NewUser(identifier, *credentials)

	usersList := append(utils.GetCredentials().Users, *user)

	data := map[string][]models.User{
		"users": usersList,
	}

	JSONData, _ := json.MarshalIndent(data, "", "    ")
	utils.WriteStringInFile(fmt.Sprintf("%s/git-credentials.json", utils.GetHomeDir()), string(JSONData))
}