package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()

	for {
		fmt.Print("Pokedex > ")
		ok := scanner.Scan()

		if ok {
			input := scanner.Text()
			data := cleanInput(input)

			if len(data) > 0 {
				command := data[0]

				if cm, ok := commands[command]; ok {
					cm.callback()
				} else {
					fmt.Println("Unknown command")
				}
			}
		}
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}

func helpCommand() error {
	commands := getCommands()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")

	for code, command := range commands {
		fmt.Printf("%s: %s\n", code, command.description)
	}

	return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
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
	}
}
