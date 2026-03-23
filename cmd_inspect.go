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
		fmt.Printf("Name: %s\n", pokemonData.Name)
		fmt.Printf("Height: %d\n", pokemonData.Height)
		fmt.Printf("Weight: %d\n", pokemonData.Weight)

		fmt.Printf("Stats:\n")
		for _, stat := range pokemonData.Stats {
			fmt.Printf("\t-%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}

		fmt.Printf("Types:\n")
		for _, t := range pokemonData.Types {
			fmt.Printf("\t-%s\n", t.Type.Name,)
		}

		return nil
	}

	return errors.New("you have not caught that pokemon")
}