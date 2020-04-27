package commands

import (
	"fmt"

	"../app/models"
	"../data"
	"../utils"
)

var initInvocation = []string{"i", "init"}

var initDescription = "Initialise the program, creating a git-credentials.json file in your home directory.\nThat file holds your git credentials."

// InitCommand is the help command
var InitCommand = models.NewCommand(ExecInit, initDescription, initInvocation)

// ExecInit executes the init command
func ExecInit() {

	path := fmt.Sprintf("%s/git-credentials.json", utils.GetHomeDir())
	if !utils.FileExists(path) {
		utils.WriteStringInFile(path, data.GitCredentialsJSONStub)
		return
	}

	fmt.Println(fmt.Sprintf("%s/git-credentials.json file already exists.", utils.GetHomeDir()))
}
