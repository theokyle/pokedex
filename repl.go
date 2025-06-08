package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedex/internal/pokeapi"
	"strings"
)

type Config struct {
	pokeapiClient pokeapi.Client
	next_url      *string
	previous_url  *string
}

func cleanInput(text string) []string {
	split_text := strings.Split(text, " ")
	var clean_text []string
	for _, word := range split_text {
		if word != "" {
			clean_text = append(clean_text, strings.ToLower(strings.TrimSpace(word)))
		}
	}
	return clean_text
}

func startRepl(cfg *Config) {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		input := cleanInput(text)

		command, ok := commands[input[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		err := command.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    CommandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    CommandMapf,
		},
		"mapb": {
			name:        "map",
			description: "Get the previous page of locations",
			callback:    CommandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    CommandExit,
		},
	}
}
