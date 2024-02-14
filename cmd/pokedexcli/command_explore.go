package main

import (
	"errors"
	"fmt"
)

func callbackExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no location provided")
	}
	locationAreaName := args[0]

	resp, err := cfg.pokeApiClient.GetLocationArea(locationAreaName)

	if err != nil {
		return err
	}

	fmt.Printf("%s pokemons found:\n", locationAreaName)
	for _, pokemon := range resp.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}

	cfg.totalExperience += 1
	fmt.Printf("By exploring you gained 1xp, total: %v\n", cfg.totalExperience)

	return nil
}
