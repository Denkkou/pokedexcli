package main

import (
	"fmt"
	"errors"
)

// Retrieve from pokedex stat info if pokemon is logged
func commandPokedex(config *config, args string) error {
	// If pokedex is empty, exit
	if len(config.pokedex) == 0 {
		return errors.New("there are no pokemon in your pokedex")
	}

	fmt.Printf("Your Pokedex:\n")
	for pokemon, _ := range config.pokedex {
		fmt.Printf(" - %s\n", pokemon)
	}

	return nil
}