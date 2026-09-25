package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()

	c := &config{
		commands: getCommands(),
	}

	for {
		fmt.Print("Pokedex > ")
		ok := scanner.Scan()

		if ok {
			input := scanner.Text()
			data := cleanInput(input)

			if len(data) > 0 {
				command := data[0]

				if cm, ok := commands[command]; ok {
					cm.callback(c)
				} else {
					fmt.Println("Unknown command")
				}
			}
		}
	}
}
