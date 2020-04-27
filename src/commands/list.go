package commands

import (
	"fmt"

	"../app/models"
	"../utils"
)

var listInvocation = []string{"l", "list"}

var listDescription = "Lists the users available from the git-credentials.json file."

// ListCommand is the help command
var ListCommand = models.NewCommand(ExecList, listDescription, listInvocation)

// ExecList prints the available credentials
func ExecList() {

	credentialsPath := fmt.Sprintf("%s/git-credentials.json", utils.GetHomeDir())

	if utils.FileExists(credentialsPath) {
		users := utils.GetCredentials()

		fmt.Println(users.List())
		return
	}
	fmt.Println("git-credentials.json file does not exists.\nGenerate it with the i (init) command.")
}
