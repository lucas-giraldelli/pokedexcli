package main

import (
	"time"

	"github.com/lucas-giraldelli/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeApiClient        pokeapi.Client
	nextLocationArea     *string
	previousLocationArea *string
}

func main() {
	interval := time.Minute * 5
	cfg := config{
		pokeApiClient: pokeapi.NewClient(interval),
	}

	startRepl(&cfg)
}
