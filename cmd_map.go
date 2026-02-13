package main

import (
	"fmt"
	"errors"
)

func commandMap(config *config, parameter string) error {
	// Get the map data from the client's GetLocations function
	mapDataResponse, err := config.client.GetLocations(config.next)
	if err != nil {
		return err
	}

	// Update the next and previous URLs using the response
	config.next = mapDataResponse.Next
	config.prev = mapDataResponse.Previous

	// Print out location results
	for _, location := range mapDataResponse.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapBack(config *config, parameter string) error {
	if config.prev == nil {
		return errors.New("Already on first page.")
	}

	mapDataResponse, err := config.client.GetLocations(config.prev)
	if err != nil {
		return err
	}

	config.next = mapDataResponse.Next
	config.prev = mapDataResponse.Previous

	for _, location := range mapDataResponse.Results {
		fmt.Println(location.Name)
	}
	
	return nil
}