package main

import "github.com/roxensox/pokedexcli/internal/pokeapi"

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
	}
}

type config struct {
	pokeapiClient pokeapi.Client
	Next          *string
	Previous      *string
}
