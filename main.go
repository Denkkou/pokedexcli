package main

import (
	"time"
	"github.com/Denkkou/pokedexcli/internal/pokeapi"
)

func main() {
	// Create a new instance of a HTTP client and
	// pass it to the REPL via the config struct
	client := pokeapi.NewClient(5 * time.Second)
	config := &config {
		client: client,
	}

	startRepl(config)
}