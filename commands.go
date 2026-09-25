package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
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
	}
}

func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}

func helpCommand(c *config) error {
	commands := getCommands()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")

	for code, command := range commands {
		fmt.Printf("%s: %s\n", code, command.description)
	}

	return nil
}

type apiResp[T any] struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []T
}

type area struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func mapDirectional(c *config, isForward bool) error {
	currentURL := "https://pokeapi.co/api/v2/location-area/"
	var body []byte

	if isForward {
		if c.Next != "" {
			currentURL = c.Next
		}
	} else {
		if c.Previous != "" {
			currentURL = c.Previous
		}
	}

	cachedData, ok := c.cache.Get(currentURL)

	if ok {
		body = cachedData
	} else {
		res, err := http.Get(currentURL)
		if err != nil {
			return err
		}

		defer res.Body.Close()

		body2, err := io.ReadAll(res.Body)

		body = body2
		if err != nil {
			return err
		}

		c.cache.Add(currentURL, body)
	}

	var data apiResp[area]

	if err := json.Unmarshal(body, &data); err != nil {
		return err
	}

	for _, v := range data.Results {
		fmt.Println(v.Name)
	}

	c.Next = data.Next
	c.Previous = data.Previous

	return nil
}

func mapCommand(c *config) error {
	return mapDirectional(c, true)
}

func mapBCommand(c *config) error {
	return mapDirectional(c, false)
}
