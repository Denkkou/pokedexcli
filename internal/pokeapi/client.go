package pokeapi

import (
	"net/http"
	"time"
	"encoding/json"
	"io"
)

// The HTTP client to be used throughout the program
type Client struct {
	httpClient http.Client
}

func NewClient(timeout time.Duration) Client {
	c := Client {
		httpClient: http.Client {
			Timeout: timeout,
		},
	}

	return c
}

// This function is part of the Client struct so we can
// use it wherever a client is used, eg, the map commands.
func (c *Client) GetLocations(url *string) (MapData, error) {
	// Lane's implementation has this function take a URL
	// and perform all of the HTTP requests, returning
	// the populated MapData struct.
	// The map functions then only need to call this function
	// to retrieve the data they need, keeping the API interactions
	// localised to this package.

	// Build the starting URL, unless we are passed a valid
	// URL from the caller (eg, the Next page)
	fullURL := baseURL + "/location-area"
	if url != nil {
		fullURL = *url
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

	mapDataResponse := MapData{}
	err = json.Unmarshal(data, &mapDataResponse) 
	if err != nil {
		return MapData{}, err
	}

	return mapDataResponse, nil
}