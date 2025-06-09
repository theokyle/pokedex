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
	pokedex       map[string]Pokemon
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

		if len(input) > 1 {
			paramater := input[1]
			err := command.callback(cfg, paramater)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			err := command.callback(cfg, "")
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
