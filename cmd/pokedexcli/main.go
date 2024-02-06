package main

import "github.com/lucas-giraldelli/pokedexcli/internal/pokeapi"

type config struct {
	pokeApiClient        pokeapi.Client
	nextLocationArea     *string
	previousLocationArea *string
}

func main() {
	cfg := config{
		pokeApiClient: pokeapi.NewClient(),
	}

	startRepl(&cfg)
}
