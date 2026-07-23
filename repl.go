package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	fmt.Println("Welcome to the Pokedex!")
	reader := bufio.NewScanner(os.Stdin)
	config := config{
		previous: 1,
		next:     1,
	}
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		if command, ok := getCommands()[commandName]; ok {
			err := command.callback(&config)
			if err != nil {
				fmt.Printf("Error executing command '%s': %v\n", commandName, err)
			}
		} else {
			fmt.Printf("Unknown command: '%s'\n", commandName)
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(config *config) error
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
			description: "Displays next 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays back 20 location areas",
			callback:    commandMapB,
		},
	}
}

type config struct {
	previous int
	next     int
}
