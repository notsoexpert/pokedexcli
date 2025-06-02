package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			text := scanner.Text()
			cleanedInput := cleanInput(text)
			if len(cleanedInput) > 0 {
				fmt.Printf("Your command was: %s\n", cleanedInput[0])
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
