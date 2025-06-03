package commands

import (
	"fmt"
	"os"
	"net/http"
	"log"
	"io"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

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
	cliCommands["map"] = cliCommand{
		name:			"map",
		description:	"Displays the names of 20 locations",
		callback:		commandMap,
	}
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

func commandMap() error {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))

	return nil
}