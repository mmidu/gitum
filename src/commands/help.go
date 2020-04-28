package commands

import (
	"../app/models"
)

var helpInvocation = []string{"h", "help"}

var helpDescription = "Prints the program description and the available commands."

// HelpCommand is the help command
var HelpCommand = models.NewCommand(ExecHelp, helpDescription, helpInvocation)

// commands holds the commands values and functionalities
var commands = map[string]*models.Command{
	"init":    InitCommand,
	"import":  ImportCommand,
	"current": CurrentCommand,
	"list":    ListCommand,
	"set":     SetCommand,
	"add":     AddCommand,
	"remove":  RemoveCommand,
}

// ExecHelp prints the helper text
func ExecHelp() {
	help := models.NewCommand(ExecHelp, helpDescription, helpInvocation)
	for _, command := range commands {
		command.PrintHelper()
	}
	help.PrintHelper()
}
