package commands

import (
	"encoding/json"
	"fmt"
	"os"

	"../app/models"
	"../utils"
)

var removeInvocation = []string{"r", "remove"}

var removeDescription = "Removes a user."

// RemoveCommand is the remove command
var RemoveCommand = models.NewCommand(ExecRemove, removeDescription, removeInvocation)

// ExecRemove executes the remove command
func ExecRemove() {
	if len(os.Args) > 3 && (os.Args[2] == "-i" || os.Args[2] == "--identifier") {
		users := utils.GetCredentials()
		userIdentifier := os.Args[3]

		var usersList []models.User //append(utils.GetCredentials().Users, *user)

		for _, user := range users.Users {
			if user.Identifier != userIdentifier {
				credentials := models.NewCredentials(user.Credentials.Username, user.Credentials.Password, user.Credentials.Domain)
				user := models.NewUser(user.Identifier, *credentials)
				usersList = append(usersList, *user)
			}
		}
		data := map[string][]models.User{
			"users": usersList,
		}

		JSONData, _ := json.MarshalIndent(data, "", "    ")
		utils.WriteStringInFile(fmt.Sprintf("%s/git-credentials.json", utils.GetHomeDir()), string(JSONData))
		return
	}
	fmt.Println("Indicate the user you want to remove adding the identifier after the flag -i (--identifier)")
}
