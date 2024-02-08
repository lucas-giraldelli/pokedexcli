package main

import (
	"errors"
	"fmt"
)

func callbackInspect(cfg *config, args ...string) error {
	if (len(args) != 1) {
		return errors.New("no pokemon specified")
	}

	pokemonName := args[0]

	pokemon, ok := cfg.caughtPokemon[pokemonName]

	if !ok {
		return errors.New("pokemon not caught yet")
	}

	fmt.Printf("Name %s\n", pokemon.Name)
	fmt.Printf("Height %v\n", pokemon.Height)
	fmt.Printf("Weight %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, s := range pokemon.Stats {
		fmt.Printf("%s - %v\n", s.Stat.Name, s.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf(" - %s\n", t.Type.Name)
	}

	return nil
}
