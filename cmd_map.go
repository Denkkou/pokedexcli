package main

import (
	"fmt"
	"io"
	"net/http"
	"encoding/json"
)

type MapData struct {
	Count int `json:"count"`
	Next *string `json:"next"`
	Previous *string `json:"previous"`
	Results []struct {
		Name string `json:"name"`
		URL string `json:"url"`
	} `json:"results"`
}

func commandMap(config *config) error {
	// Make sure url is either start, or our next
	baseUrl := "https://pokeapi.co/api/v2/location-area/"
	url := config.next
	if url == nil {
		url = &baseUrl
	}

	// Make GET request for next 20 location areas
	res, err := http.Get(*url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// Unmarshal
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	mapData := MapData{}
	if err := json.Unmarshal(data, &mapData); err != nil {
		return err
	}

	// Set next and previous
	config.prev = mapData.Previous
	config.next = mapData.Next

	// Print out location results
	for _, location := range mapData.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapBack(config *config) error {
	if config.prev == nil {
		return fmt.Errorf("You are already on the first page!")
	}

	// Make request but for previous url
	url := config.prev
	res, err := http.Get(*url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// Unmarshal
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	mapData := MapData{}
	if err := json.Unmarshal(data, &mapData); err != nil {
		return err
	}

	// Set next and previous
	config.prev = mapData.Previous
	config.next = mapData.Next

	// Print out location results
	for _, location := range mapData.Results {
		fmt.Println(location.Name)
	}

	return nil
}