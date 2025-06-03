package main

import (
	"bufio"
	"fmt"
	"strings"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			text := scanner.Text()
			cleanedInput := cleanInput(text)
			if len(cleanedInput) == 0 {
				continue
			}

			command, ok := cliCommands[cleanedInput[0]]
			if !ok {
				fmt.Println("Unknown command")
				continue
			}

			err := command.callback()
			if err != nil {
				fmt.Printf("%s", err)
			}

		}
	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	text = strings.TrimSpace(text)
	splitInput := strings.Split(text, " ")
	cleanedInput := []string{}
	for _, str := range splitInput {
		if str != "" {
			cleanedInput = append(cleanedInput, str)
		}
	}
	return cleanedInput
}