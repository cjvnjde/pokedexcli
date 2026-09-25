package main

import (
	"strings"
)

func cleanInput(text string) []string {
	splitString := strings.Split(strings.ToLower(text), " ")
	var cleanedSplitString []string = make([]string, 0, len(splitString))

	for _, v := range splitString {
		if v != "" {
			cleanedSplitString = append(cleanedSplitString, v)
		}
	}

	return cleanedSplitString
}
