package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    helpCommand,
		},
		"map": {
			name:        "map",
			description: "Displays the name of next 20 location areas",
			callback:    mapCommand,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the name of prevous 20 location areas",
			callback:    mapBCommand,
		},
		"explore": {
			name:        "explore",
			description: "Displays pokemons in the location",
			callback:    exploreCommand,
		},
	}
}

func commandExit(c *config, param string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}

func helpCommand(c *config, param string) error {
	commands := getCommands()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")

	for code, command := range commands {
		fmt.Printf("%s: %s\n", code, command.description)
	}

	return nil
}
