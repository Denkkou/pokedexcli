package pokeapi

import (
	"net/http"
	"encoding/json"
	"io"
)

// Retrieve from HTTP or cache ExploreData pertaining to given areaName
func (c *Client) Explore(areaName string) (ExploreData, error) {
	fullURL := baseURL + "/location-area/" + areaName

	// Check cache for an entry
	entry, ok := c.cache.Get(fullURL)
	if ok {
		exploreDataResponse := ExploreData{}
		err := json.Unmarshal(entry, &exploreDataResponse)
		if err != nil {
			return ExploreData{}, err
		}
		return exploreDataResponse, nil
	}

	// Build request
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return ExploreData{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return ExploreData{}, err
	}
	defer res.Body.Close()

	// Read
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return ExploreData{}, err
	}

	// Cache data
	c.cache.Add(fullURL, data)

	// Unmarshal
	exploreDataResponse := ExploreData{}
	err = json.Unmarshal(data, &exploreDataResponse)
	if err != nil {
		return ExploreData{}, err
	}

	return exploreDataResponse, nil
}