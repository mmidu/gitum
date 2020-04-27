package routes

import (
	"../app/models"
	"../commands"
)

// List holds the commands keys
var List = map[string]string{

	"i":    "init",
	"init": "init",

	"s":   "set",
	"set": "set",

	"c":       "current",
	"current": "current",

	"ui":     "import",
	"import": "import",

	"l":    "list",
	"list": "list",

	"h":    "help",
	"help": "help",

	"a":   "add",
	"add": "add",

	"r":      "remove",
	"remove": "remove",
}

// Map holds the commands values and functionalities
var Map = map[string]*models.Command{
	"help":    commands.HelpCommand,
	"init":    commands.InitCommand,
	"import":  commands.ImportCommand,
	"current": commands.CurrentCommand,
	"list":    commands.ListCommand,
	"set":     commands.SetCommand,
	"add":     commands.AddCommand,
	"remove":  commands.RemoveCommand,
}
