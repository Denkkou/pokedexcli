package main

import (
	"fmt"
	"errors"
)

func commandCatch(config *config, pokemonName string) error {
	if pokemonName == "" {
		return errors.New("Pokemon name invalid.")
	}

	// Do the catch request to get information on the pokemon
	catchResponse, err := config.client.Catch(pokemonName)
	if err != nil {
		return err
	}
	
	// Then, take its base experience to do the random chance
	fmt.Printf("Throwing a Pokeball at %s...\n", catchResponse.Name)
	//...

	// Then add to pokedex
	return nil
}