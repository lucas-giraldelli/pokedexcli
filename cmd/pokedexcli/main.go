package main

import (
	"time"

	"github.com/lucas-giraldelli/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeApiClient					pokeapi.Client
	nextLocationArea			*string
	previousLocationArea	*string
	caughtPokemon 				map[string]pokeapi.Pokemon
	totalExperience				int
}

func main() {
	interval := time.Hour
	cfg := config{
		pokeApiClient: pokeapi.NewClient(interval),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
		totalExperience: 0,
	}

	startRepl(&cfg)
}
