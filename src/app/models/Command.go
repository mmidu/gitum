package models

import "fmt"

// Command is the command struct
type Command struct {
	Function    func()
	Description string
	Invocation  []string
}

// NewCommand makes a new command
func NewCommand(function func(), description string, invocation []string) *Command {
	cmd := Command{
		Function:    function,
		Description: description,
		Invocation:  invocation,
	}
	return &cmd
}

// PrintDescription prints the description string of the command
func (c Command) PrintDescription() {
	fmt.Println(c.Description)
}

// Exec executes the command
func (c Command) Exec() {
	c.Function()
}

// PrintHelper prints the command helper
func (c Command) PrintHelper() {
	helper := fmt.Sprintf("%s: %s", c.Invocation, c.Description)
	fmt.Println(helper)
}
