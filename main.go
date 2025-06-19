package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/notsoexpert/pokedexcli/internal/pokeapi"
	"github.com/notsoexpert/pokedexcli/internal/pokecommand"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var location pokeapi.Location
	location.Base = "https://pokeapi.co/api/v2/location-area/"

	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			text := scanner.Text()
			cleanedInput := cleanInput(text)
			if len(cleanedInput) == 0 {
				continue
			}
			cleanedInput = append(cleanedInput, "")

			command, ok := pokecommand.Execute(cleanedInput[0])
			if !ok {
				fmt.Println("Unknown command")
				continue
			}

			err := command.Callback(&location, cleanedInput[1])
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
