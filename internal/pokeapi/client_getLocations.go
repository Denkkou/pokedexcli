package pokeapi

import (
	"net/http"
	"encoding/json"
	"io"
)

// Retrieve from HTTP or cache MapData pertaining to given URL
func (c *Client) GetLocations(url *string) (MapData, error) {
	// Build the starting URL, unless we are passed a valid
	// URL from the caller (eg, the Next page)
	fullURL := baseURL + "/location-area"
	if url != nil {
		fullURL = *url
	}

	// Unmarshal and return entry if it exists in cache
	entry, ok := c.cache.Get(fullURL)
	if ok {
		mapDataResponse := MapData{}
		err := json.Unmarshal(entry, &mapDataResponse) 
		if err != nil {
			return MapData{}, err
		}

		return mapDataResponse, nil
	}

	// Build request
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return MapData{}, err
	}

	res, err := c.httpClient.Do(req) 
	if err != nil {
		return MapData{}, err
	}
	defer res.Body.Close()

	// Unmarshal the data into the MapData struct
	data, err := io.ReadAll(res.Body) 
	if err != nil {
		return MapData{}, err
	}

	// Data is not cached, so cache it
	c.cache.Add(fullURL, data)

	mapDataResponse := MapData{}
	err = json.Unmarshal(data, &mapDataResponse) 
	if err != nil {
		return MapData{}, err
	}

	return mapDataResponse, nil
}