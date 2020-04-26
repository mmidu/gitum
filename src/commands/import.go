package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"../app/models"
	"../utils"
)

var importInvocation = []string{"iu", "import"}

var importDescription = "Imports the currently active user in the git-credentials.json file, if it is not already registered."

// ImportCommand is the help command
var ImportCommand = models.NewCommand(ExecImport, importDescription, importInvocation)

// ExecImport imports an existing user
func ExecImport() {
	fmt.Println("import executed")

	credentialsPath := fmt.Sprintf("%s/git-credentials.json", utils.GetHomeDir())

	if utils.FileExists(credentialsPath) {
		users := utils.GetCredentials()
		currentUser := utils.GetCurrentUser()

		for _, user := range users.Users {
			if user.Credentials.Username == currentUser["username"] &&
				user.Credentials.Password == currentUser["password"] &&
				user.Credentials.Domain == currentUser["domain"] {
				os.Exit(0)
			}
		}

		var newUser models.User

		newUser.Identifier = currentUser["username"] + "@" + strings.SplitN(currentUser["domain"], ".", -1)[0]
		newUser.Credentials.Username = currentUser["username"]
		newUser.Credentials.Password = currentUser["password"]
		newUser.Credentials.Domain = currentUser["domain"]

		users.Users = append(users.Users, newUser)

		data := map[string][]models.User{
			"users": users.Users,
		}

		JSONData, _ := json.MarshalIndent(data, "", "    ")
		utils.WriteStringInFile(fmt.Sprintf("%s/git-credentials.json", utils.GetHomeDir()), string(JSONData))
		os.Exit(0)
	} else {
		fmt.Println("git-credentials.json file does not exists.\nGenerate it with the -i (--init) flag.")
	}
	os.Exit(0)
}
