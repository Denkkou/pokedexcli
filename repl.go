package main

import (
	"fmt"
	"bufio"
	"strings"
	"os"
)

// Run the REPL with this function
func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		
		// Prepare user input
		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}
		
		// Handle commands
		if cmd, ok := getCommands()[input[0]]; ok {
			// Capture callback func's error
			err := cmd.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command\n")
			continue
		}
	}
}

func cleanInput(text string) []string {
	cleanedTextSlice := strings.Fields(strings.ToLower(text))
	return cleanedTextSlice
}

// Each cli command follows this structure
type cliCommand struct {
	name string
	description string
	callback func() error
}

func getCommands() map[string]cliCommand {
	// Define each cli command in this map
	return map[string]cliCommand {
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
	}
}