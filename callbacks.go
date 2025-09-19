package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"slices"
)

func commandExit(*config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(*config) error {
	cmds := getCommands()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	for _, c := range cmds {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}
	return nil
}

func commandMap(cfg *config) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.Next)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	cfg.Next = resp.Next
	cfg.Previous = resp.Previous
	return nil
}

func commandMapb(cfg *config) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.Previous)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	cfg.Next = resp.Next
	cfg.Previous = resp.Previous
	return nil
}

func commandExplore(cfg *config) error {
	// Displays a list of the pokemon in a specified location

	// Ensures the user supplied at least two commands
	if len(cfg.Commands) < 2 {
		return fmt.Errorf("Must enter a location name")
	}

	// Requests the location-area information
	resp, err := cfg.pokeapiClient.ListEncounters(cfg.Commands[1])

	// Checks for errors
	if err != nil {
		fmt.Printf("Error: Unable to find %s\n", cfg.Commands[1])
		return err
	}

	// Outputs list of pokemon
	fmt.Println("Found Pokemon:")
	for _, v := range resp.PokemonEncounters {
		fmt.Printf(" - %s\n", v.Pokemon.Name)
	}

	return nil
}

func commandCatch(cfg *config) error {
	if len(cfg.Commands) < 2 {
		return fmt.Errorf("Must enter a pokemon ID or name")
	}

	resp, err := cfg.pokeapiClient.GetPokemon(cfg.Commands[1])
	if err != nil {
		fmt.Printf("Error: Unable to find %s\n", cfg.Commands[1])
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", cfg.Commands[1])

	maxExp := 608.0
	minExp := 0.0
	actualExp := float64(resp.BaseExperience)
	minChance := 0.01
	maxChance := 0.85

	t := (actualExp - minExp) / (maxExp - minExp)
	p := maxChance + (minChance-maxChance)*t
	r := rand.Intn(100)

	if r < int(p*100) {
		fmt.Printf("%s was caught!\n", cfg.Commands[1])
		cfg.Caught = append(cfg.Caught, cfg.Commands[1])
	} else {
		fmt.Printf("%s escaped!\n", cfg.Commands[1])
	}
	return nil
}

func commandInspect(cfg *config) error {
	if len(cfg.Commands) < 2 {
		return fmt.Errorf("Must enter a pokemon ID or name")
	}

	if !slices.Contains(cfg.Caught, cfg.Commands[1]) {
		fmt.Printf("You have not caught a %s\n", cfg.Commands[1])
		return fmt.Errorf("You have not caught a %s", cfg.Commands[1])
	}

	resp, err := cfg.pokeapiClient.GetPokemon(cfg.Commands[1])
	if err != nil {
		fmt.Printf("Error: Unable to find %s\n", cfg.Commands[1])
		return err
	}

	fmt.Printf("Name: %s\n", resp.Name)
	fmt.Printf("Height: %v\n", resp.Height)
	fmt.Printf("Weight: %v\n", resp.Weight)
	fmt.Print("Stats:\n")
	for _, v := range resp.Stats {
		fmt.Printf(" -%s: %v\n", v.Stat.Name, v.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range resp.Types {
		fmt.Printf(" -%s\n", t.Type.Name)
	}

	return nil
}

func commandPokedex(cfg *config) error {
	if len(cfg.Caught) < 1 {
		return fmt.Errorf("No Pokemon caught")
	}
	fmt.Println("Your Pokedex:")
	for _, c := range cfg.Caught {
		fmt.Printf(" - %s\n", c)
	}
	return nil
}
