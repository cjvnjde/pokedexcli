package main

type config struct {
	commands map[string]cliCommand
	Next     string
	Previous string
}
