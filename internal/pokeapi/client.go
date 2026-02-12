package pokeapi

import (
	"net/http"
	"time"
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