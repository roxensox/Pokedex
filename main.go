package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	var command string

	commands["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	}
	commands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	}

	for {
		fmt.Print("Pokedex > ")
		input.Scan()
		input := cleanInput(input.Text())

		if len(input) > 0 {
			command = input[0]
		} else {
			command = ""
		}

		if val, ok := commands[command]; ok {
			err := val.callback()
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!\n\nUsage:")
	for _, v := range commands {
		fmt.Printf("\t%s: %s\n", v.name, v.description)
	}
	return nil
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
