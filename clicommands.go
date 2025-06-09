package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

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
	} else {
		fmt.Println("You missed!")
	}

	return nil
}
