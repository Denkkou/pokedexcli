package pokeapi

import (
	"net/http"
	"time"
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