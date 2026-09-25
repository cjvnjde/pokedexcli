package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/cjvnjde/pokedexcli/internal"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()

	c := &config{
		commands: getCommands(),
		cache:    internal.NewCache(5 * time.Second),
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
