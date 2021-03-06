package commands

import (
	"fmt"

	"../app/models"
	"../utils"
)

var currentInvocation = []string{"c", "current"}

var currentDescription = "Outputs the currently active git user."

// CurrentCommand is the help command
var CurrentCommand = models.NewCommand(ExecCurrent, currentDescription, currentInvocation)

// ExecCurrent handles the current user command
func ExecCurrent() {

	users := utils.GetCredentials()
	currentUser := utils.GetCurrentUser()

	for _, user := range users.Users {
		if user.Credentials.Username == currentUser["username"] &&
			user.Credentials.Password == currentUser["password"] &&
			user.Credentials.Domain == currentUser["domain"] {
			fmt.Println(user.Identifier)
		}
	}
}
