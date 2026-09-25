package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type area struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func location(c *config, isForward bool) error {
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
