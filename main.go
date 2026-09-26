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
		pokemons: map[string]Pokemon{},
	}

	for {
		fmt.Print("Pokedex > ")
		ok := scanner.Scan()
		err := scanner.Err()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if ok {
			input := scanner.Text()
			data := cleanInput(input)
			var command, param string

			if len(data) > 0 {
				command = data[0]
			}
			if len(data) > 1 {
				param = data[1]
			}

			if cm, ok := commands[command]; ok {
				if param != "" {
					if error := cm.callback(c, param); error != nil {
						fmt.Println(error)
					}
				} else {
					if error := cm.callback(c, ""); error != nil {
						fmt.Println(error)
					}
				}
			} else {
				fmt.Println("Unknown command")
			}
		} else {
			os.Exit(0)
		}
	}
}
