package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var Commands map[string]cliCommand

func CommandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func CommandHelp() error {
	command_list := ""
	for _, command := range Commands {
		command_list += fmt.Sprintf("%s: %s\n", command.name, command.description)
	}
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n%s", command_list)

	return nil
}
