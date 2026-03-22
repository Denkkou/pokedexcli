package main

import (
	"fmt"
	"errors"
)

// Retrieve from pokedex stat info if pokemon is logged
func commandInspect(config *config, pokemonName string) error {
	// Check if pokemon is in pokedex
	pokemonData, ok := config.pokedex[pokemonName]
	if ok {
		// Parse stat info and print
		// name, height, weight, stats, types
		// 		hp, atk, def, spatk, spdef, speed
		fmt.Printf("Name: %s\n", pokemonData.Name)

		return nil
	}

	return errors.New("you have not caught that pokemon")
}