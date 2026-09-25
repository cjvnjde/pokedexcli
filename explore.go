package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type region struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type language struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type name struct {
	Name     string   `json:"name"`
	Language language `json:"language"`
}
type pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type pokemonEncounter struct {
	Pokemon pokemon `json:"pokemon"`
}
type areaData struct {
	ID                int                `json:"id"`
	Name              string             `json:"name"`
	Region            region             `json:"region"`
	Names             []name             `json:"names"`
	PokemonEncounters []pokemonEncounter `json:"pokemon_encounters"`
}

func exploreCommand(c *config, name string) error {
	currentURL := "https://pokeapi.co/api/v2/location-area/" + name
	var body []byte

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

	var data areaData

	if err := json.Unmarshal(body, &data); err != nil {
		return err
	}

	for _, v := range data.PokemonEncounters {
		fmt.Println(v.Pokemon.Name)
	}

	return nil
}
