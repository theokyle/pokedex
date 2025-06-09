package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, string) error
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
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    CommandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    CommandExit,
		},
		"explore": {
			name:        "explore",
			description: "see pokemon located in an area",
			callback:    CommandExplore,
		},
		"catch": {
			name:        "catch",
			description: "catch a pokemon",
			callback:    CommandCatch,
		},
		"pokedex": {
			name:        "pokedex",
			description: "display list of caught pokemon",
			callback:    CommandPokedex,
		},
		"inspect": {
			name:        "inspect",
			description: "get info on a caught pokemon",
			callback:    CommandInspect,
		},
	}
}

func CommandExit(cfg *Config, paramater string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func CommandHelp(cfg *Config, paramater string) error {
	command_list := ""
	commands := getCommands()
	for _, command := range commands {
		command_list += fmt.Sprintf("%s: %s\n", command.name, command.description)
	}
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n%s", command_list)

	return nil
}

func CommandMapf(cfg *Config, paramater string) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.next_url)
	if err != nil {
		return err
	}

	cfg.next_url = locationsResp.Next
	cfg.previous_url = locationsResp.Previous

	for _, location := range locationsResp.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func CommandMapb(cfg *Config, paramater string) error {
	if cfg.previous_url == nil {
		return errors.New("you're on the first page")
	}

	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.previous_url)
	if err != nil {
		return err
	}

	cfg.next_url = locationsResp.Next
	cfg.previous_url = locationsResp.Previous

	for _, location := range locationsResp.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func CommandExplore(cfg *Config, area_name string) error {
	locationResp, err := cfg.pokeapiClient.ExploreLocation(&area_name)
	if err != nil {
		return err
	}
	fmt.Println("You have found the following pokemon:")
	for _, encounter := range locationResp.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}
	return nil
}

func CommandCatch(cfg *Config, pokemon string) error {
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)
	pokemonResp, err := cfg.pokeapiClient.PokemonDetail(&pokemon)
	if err != nil {
		return err
	}

	number := rand.Intn(500)

	if pokemonResp.BaseExperience < number {
		fmt.Printf("%s was caught\n", pokemon)
		AddToPokedex(cfg, pokemonResp)
		fmt.Println("You may now inspect it with the inspect command")
	} else {
		fmt.Println("You missed!")
	}

	return nil
}

func CommandPokedex(cfg *Config, parameter string) error {
	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.pokedex {
		fmt.Printf(" - %s\n", pokemon.name)
	}

	return nil
}

func CommandInspect(cfg *Config, pokemon_to_inspect string) error {
	pokemon, ok := cfg.pokedex[pokemon_to_inspect]
	if !ok {
		return errors.New("Pokemon not found")
	}

	fmt.Printf("Name: %s\n", pokemon.name)
	fmt.Printf("Height: %v\n", pokemon.height)
	fmt.Printf("Weight: %v\n", pokemon.weight)
	fmt.Println("Stats:")
	for key, value := range pokemon.stats {
		fmt.Printf("  -%s: %v\n", key, value)
	}
	fmt.Println("Types:")
	for _, ptype := range pokemon.types {
		fmt.Printf("  - %s\n", ptype)
	}

	return nil
}
