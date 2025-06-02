package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var cliCommands map[string]cliCommand

func init() {
	cliCommands = make(map[string]cliCommand)
	cliCommands["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	}
	cliCommands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	}
}

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

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for _, cliCommand := range cliCommands {
		fmt.Printf("%s: %s\n", cliCommand.name, cliCommand.description)
	}
	return nil
}
