package pokecommand

import (
	"fmt"
	"os"

	"github.com/notsoexpert/pokedexcli/internal/pokelocation"
)

type CLICommand struct {
	name        string
	description string
	Callback    func(*pokelocation.Location) error
}

var gCommands map[string]CLICommand

func init() {
	gCommands = make(map[string]CLICommand)
	gCommands["exit"] = CLICommand{
		name:        "exit",
		description: "Exit the Pokedex",
		Callback:    commandExit,
	}
	gCommands["help"] = CLICommand{
		name:        "help",
		description: "Displays a help message",
		Callback:    commandHelp,
	}
	gCommands["map"] = CLICommand{
		name:        "map",
		description: "Displays the names of 20 locations procedurally",
		Callback:    commandMap,
	}
	gCommands["mapb"] = CLICommand{
		name:        "mapb",
		description: "Displays the names of the previous 20 locations from 'map'",
		Callback:    commandMapB,
	}
}

func commandExit(location *pokelocation.Location) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(location *pokelocation.Location) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for _, cliCommand := range gCommands {
		fmt.Printf("%s: %s\n", cliCommand.name, cliCommand.description)
	}
	return nil
}

func commandMap(location *pokelocation.Location) error {
	var url string

	if location.Next == nil {
		if location.Previous == nil {
			url = location.Current
		} else {
			fmt.Printf("You're on the last page.\n")
			return nil
		}
	} else {
		url = *location.Next
	}

	err := pokelocation.UpdateLocations(url, location)
	if err != nil {
		return err
	}

	location.Current = url
	pokelocation.PrintLocations(location)
	return nil
}

func commandMapB(location *pokelocation.Location) error {
	var url string

	if location.Previous == nil {
		fmt.Printf("You're on the first page.\n")
		return nil
	} else {
		url = *location.Previous
	}

	err := pokelocation.UpdateLocations(url, location)
	if err != nil {
		return err
	}

	location.Current = url
	pokelocation.PrintLocations(location)
	return nil
}

func Execute(key string) (CLICommand, bool) {
	cmd, ok := gCommands[key]
	return cmd, ok
}
