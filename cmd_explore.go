package main

import (
	"fmt"
	"errors"
)

func commandExplore(config *config, areaName string) error {
	if areaName == "" {
		return errors.New("Area name invalid.")
	}

	// Get explore data
	exploreResponse, err := config.client.Explore(areaName)
	if err != nil {
		return err
	}

	// Print results
	fmt.Printf("Exploring %s...\n", exploreResponse.Name)
	fmt.Println("Found Pokemon: ")
	for _, encounter := range exploreResponse.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	return nil
}