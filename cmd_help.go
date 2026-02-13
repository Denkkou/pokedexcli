package main
import (
	"fmt"
)

func commandHelp(config *config) error {
	fmt.Println("\nWelcome to the Pokedex!")

	// Generate a usage section by iterating command registry
	fmt.Println("Usage:")
	fmt.Println("")
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	
	return nil
}