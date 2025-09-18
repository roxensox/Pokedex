package main

import (
	"github.com/roxensox/pokedexcli/internal/pokeapi"
)

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(),
	}

	startRepl(&cfg)
}
