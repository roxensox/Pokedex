package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// TEST
	/*
		REPL for the Pokedex
	*/
	// Creates a new scanner object for input
	input := bufio.NewScanner(os.Stdin)
	// Creates a new string variable to hold the command
	var command string

	// Registers the exit command
	commands["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	}
	// Registers the help command
	commands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	}

	for {
		// Prints prompt
		fmt.Print("Pokedex > ")
		// Collects input
		input.Scan()
		// Cleans the input
		input := cleanInput(input.Text())

		// If input was provided, extracts the command
		if len(input) > 0 {
			command = input[0]
		} else {
			// Otherwise, sets the command to an empty string
			command = ""
		}

		// If the command exists, the callback is called
		if val, ok := commands[command]; ok {
			err := val.callback()
			// Any erros are printed
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		} else {
			// Prints unknown command message if command doesn't exist
			fmt.Println("Unknown command")
		}
	}
}

func commandExit() error {
	/*
		Exits the program
	*/
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	/*
		Prints the help message
	*/
	fmt.Println("Welcome to the Pokedex!\n\nUsage:")
	// Loops through all registered commands to print details
	for _, v := range commands {
		fmt.Printf("\t%s: %s\n", v.name, v.description)
	}
	return nil
}

func cleanInput(text string) []string {
	// Separates all input terms into a slice of strings
	return strings.Fields(strings.ToLower(text))
}
