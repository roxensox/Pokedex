package main

import (
	"fmt"
	"log"
	"os"
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
