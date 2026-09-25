package main

import (
	"fmt"
)

func pokedexCommand(c *config, name string) error {
	if len(c.pokemons) == 0 {
		fmt.Println("You haven't caught any pokemon yet")
	}

	fmt.Println("Your Pokedex")
	for _, pokemon := range c.pokemons {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}
