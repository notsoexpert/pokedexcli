package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Location) error
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
		description:	"Displays the names of 20 locations procedurally",
		callback:		commandMap,
	}
	cliCommands["mapb"] = cliCommand{
		name:			"mapb",
		description:	"Displays the names of the previous 20 locations from 'map'",
		callback:		commandMapB,
	}
}

func commandExit(location *Location) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(location *Location) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for _, cliCommand := range cliCommands {
		fmt.Printf("%s: %s\n", cliCommand.name, cliCommand.description)
	}
	return nil
}

func commandMap(location *Location) error {
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

	err := updateLocations(url, location)
	if err != nil {
		return err
	}

	location.Current = url
	printLocations(location)
	return nil
}

func commandMapB(location *Location) error {
	var url string

	if location.Previous == nil {
		fmt.Printf("You're on the first page.\n")
		return nil
	} else {
		url = *location.Previous
	}

	err := updateLocations(url, location)
	if err != nil {
		return err
	}
	
	location.Current = url
	printLocations(location)
	return nil
}