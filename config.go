package main

import "github.com/cjvnjde/pokedexcli/internal"

type config struct {
	commands map[string]cliCommand
	Next     string
	Previous string
	cache    *internal.Cache
	pokemons map[string]Pokemon
}
