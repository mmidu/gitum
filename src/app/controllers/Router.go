package controllers

import (
	"fmt"
	"os"

	"../../routes"
	"../models"
)

// Router is the list of available commands
type Router struct {
	Keys map[string]string
	List map[string]*models.Command
}

// NewRouter inits the router
func NewRouter(keys map[string]string, commands map[string]*models.Command) *Router {
	router := Router{
		Keys: keys,
		List: commands,
	}
	return &router
}

// Route gets the current passed arguments
func (router Router) Route() {
	args := os.Args
	_ = args
	if len(os.Args) > 1 {
		firstArg := os.Args[1]
		cmd := router.Keys[firstArg]
		if cmd != "" {
			routes.Map[cmd].Function()
			return
		}
		fmt.Println(fmt.Sprintf("%s: command does not exist.", cmd))
	}
}

// AppRouter generates the app router
var AppRouter = NewRouter(routes.List, routes.Map)
