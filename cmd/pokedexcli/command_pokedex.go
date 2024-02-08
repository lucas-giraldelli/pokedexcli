package main

import (
	"fmt"
)

func callbackPokedex(cfg *config, args ...string) error {
	fmt.Println("Your pokedex pokemons:")
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf(" - %v\n", pokemon.Name)
	}

	return nil
}
