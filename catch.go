package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

func catchCommand(c *config, name string) error {
	currentURL := "https://pokeapi.co/api/v2/pokemon/" + name
	var body []byte
	fmt.Printf("Throwing a Pokeball at %s...\n", name)

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

	var data Pokemon

	if err := json.Unmarshal(body, &data); err != nil {
		return err
	}

	if isCought(data.BaseExperience) {
		fmt.Printf("%s was caught!\n", data.Name)
		c.pokemons[data.Name] = data
	} else {
		fmt.Printf("%s escaped!\n", data.Name)
	}

	return nil
}

func isCought(experience int) bool {
	chance := rand.Intn(100)

	return experience < chance
}
