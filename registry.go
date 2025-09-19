package main

import (
	"github.com/roxensox/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays next page of map locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous page of map locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Displays a list of all pokemon in the specified area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Tries to catch a specified pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "catch",
			description: "Displays the details of a specified pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Displays all pokemon caught by the user",
			callback:    commandPokedex,
		},
	}
}

type config struct {
	pokeapiClient pokeapi.Client
	Commands      []string
	Caught        []string
	Next          *string
	Previous      *string
}
