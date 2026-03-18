package pokeapi

import (
	"net/http"
	"encoding/json"
	"io"
)

func (c *Client) Catch(pokemonName string) (PokemonData, error) {
	fullURL := baseURL + "/pokemon/" + pokemonName

	// Check the cache for an entry
	entry, ok := c.cache.Get(fullURL)
	if ok {
		pokemonDataResponse := PokemonData{}
		err := json.Unmarshal(entry, &pokemonDataResponse)
		if err != nil {
			return PokemonData{}, err
		}
		return pokemonDataResponse, nil
	}

	// Build request
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return PokemonData{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonData{}, err
	}
	defer res.Body.Close()

	// Read
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonData{}, err
	}

	// Cache the data
	c.cache.Add(fullURL, data)

	// Unmarshal
	pokemonDataResponse := PokemonData{}
	err = json.Unmarshal(data, &pokemonDataResponse)
	if err != nil {
		return PokemonData{}, err
	}

	return pokemonDataResponse, nil
}