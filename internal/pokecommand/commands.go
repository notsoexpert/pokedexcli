package pokecommand

import (
	"fmt"
	"os"

	"github.com/notsoexpert/pokedexcli/internal/pokeapi"
)

type CLICommand struct {
	name        string
	description string
	Callback    func(*pokeapi.Location, string) error
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
	gCommands["explore"] = CLICommand{
		name:        "explore",
		description: "Displays the names of all Pokemon discoverable from the provided area",
		Callback:    commandExplore,
	}
	gCommands["catch"] = CLICommand{
		name:        "catch",
		description: "Attempt to catch the Pokemon of the provided name",
		Callback:    commandCatch,
	}
	gCommands["inspect"] = CLICommand{
		name:        "inspect",
		description: "List the details of a caught Pokemon",
		Callback:    commandInspect,
	}
	gCommands["pokedex"] = CLICommand{
		name:        "pokedex",
		description: "List the names of all caught Pokemon",
		Callback:    commandPokedex,
	}
}

func commandExit(location *pokeapi.Location, arg string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(location *pokeapi.Location, arg string) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for _, cliCommand := range gCommands {
		fmt.Printf("%s: %s\n", cliCommand.name, cliCommand.description)
	}
	return nil
}

func commandMap(location *pokeapi.Location, arg string) error {
	var url string

	if location.Next == nil {
		if location.Previous == nil {
			url = location.Base + location.Current
		} else {
			fmt.Printf("You're on the last page.\n")
			return nil
		}
	} else {
		url = *location.Next
	}

	err := pokeapi.UpdateLocations(url, location)
	if err != nil {
		return err
	}

	location.Current = url
	pokeapi.PrintLocations(location)
	return nil
}

func commandMapB(location *pokeapi.Location, arg string) error {
	var url string

	if location.Previous == nil {
		fmt.Printf("You're on the first page.\n")
		return nil
	} else {
		url = *location.Previous
	}

	err := pokeapi.UpdateLocations(url, location)
	if err != nil {
		return err
	}

	location.Current = url
	pokeapi.PrintLocations(location)
	return nil
}

func commandExplore(location *pokeapi.Location, arg string) error {
	pokeapi.PrintExploration(arg, location)
	return nil
}

func commandCatch(location *pokeapi.Location, arg string) error {
	pokeapi.AttemptCatch(arg)
	return nil
}

func commandInspect(location *pokeapi.Location, arg string) error {
	pokeapi.InspectPokemon(arg)
	return nil
}

func commandPokedex(location *pokeapi.Location, arg string) error {
	pokeapi.ListPokemon()
	return nil
}

func Execute(input string) (CLICommand, bool) {
	cmd, ok := gCommands[input]
	return cmd, ok
}
