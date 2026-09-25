package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		ok := scanner.Scan()

		if ok {
			input := scanner.Text()
			data := cleanInput(input)

			if len(data) > 0 {
				firstWord := data[0]
				fmt.Printf("Your command was: %s\n", firstWord)
			}
		}
	}
}
