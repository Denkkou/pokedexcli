package main

import (
	"fmt"
	"errors"
	"math/rand"
)

func commandCatch(config *config, pokemonToCatch string) error {
	if pokemonToCatch == "" {
		return errors.New("Pokemon name invalid.")
	}

	// Do the catch request to get information on the pokemon
	catchResponse, err := config.client.Catch(pokemonToCatch)
	if err != nil {
		return err
	}

	pokemonName := catchResponse.Name
	
	// Initiate catch
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	// Then, take its base experience to do the random chance
	baseExperience := catchResponse.BaseExperience
	const threshold = 50
	randNum := rand.Intn(baseExperience)

	if randNum > threshold {
		return fmt.Errorf("Failed to catch %s", pokemonName)		
	}

	// If caught, add to pokedex
	fmt.Printf("%s was caught!\n", pokemonName)

	fmt.Printf("Added %s's data to the Pokedex!\n", pokemonName)
	config.pokedex[pokemonName] = catchResponse


	return nil
}