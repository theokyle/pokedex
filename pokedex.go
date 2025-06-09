package main

import (
	"pokedex/internal/pokeapi"
)

type Pokemon struct {
	name   string
	height int
	weight int
	stats  map[string]int
	types  []string
}

func AddToPokedex(cfg *Config, info pokeapi.RespPokemon) error {
	newPokemon := Pokemon{
		name:   info.Name,
		height: info.Height,
		weight: info.Weight,
		stats:  make(map[string]int),
		types:  make([]string, 0),
	}

	for _, stat := range info.Stats {
		newPokemon.stats[stat.Stat.Name] = stat.BaseStat
	}

	for _, ptype := range info.Types {
		newPokemon.types = append(newPokemon.types, ptype.Type.Name)
	}

	cfg.pokedex[info.Name] = newPokemon

	return nil
}
