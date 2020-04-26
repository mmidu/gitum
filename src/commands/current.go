package commands

import (
	"fmt"
	"os"

	"../app/models"
	"../utils"
)

var currentInvocation = []string{"c", "current"}

var currentDescription = "Outputs the currently active git user."

// CurrentCommand is the help command
var CurrentCommand = models.NewCommand(ExecCurrent, currentDescription, currentInvocation)

// ExecCurrent handles the current user command
func ExecCurrent() {
	fmt.Println("current executed")

	users := utils.GetCredentials()
	currentUser := utils.GetCurrentUser()

	for _, user := range users.Users {
		if user.Credentials.Username == currentUser["username"] &&
			user.Credentials.Password == currentUser["password"] &&
			user.Credentials.Domain == currentUser["domain"] {
			fmt.Println(user.Identifier)
			os.Exit(0)
		}
	}

	os.Exit(0)
}
