package pokeapi

import (
	"net/http"
	"time"
	"encoding/json"
	"io"
	"github.com/Denkkou/pokedexcli/internal/pokecache"
)

// The HTTP client to be used throughout the program
type Client struct {
	httpClient http.Client
	cache *pokecache.Cache
}

func NewClient(timeout time.Duration) Client {
	c := Client {
		httpClient: http.Client {
			Timeout: timeout,
		},
		cache: pokecache.NewCache(timeout),
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