package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func callbackCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon provided")
	}
	pokemonName := args[0]

	pokemon, err := cfg.pokeApiClient.GetPokemon(pokemonName)

	if err != nil {
		return err
	}

	baseTreshold := 80
	threshold := baseTreshold + cfg.totalExperience
	randNum := rand.Intn(int(float64(pokemon.BaseExperience) * 1.4))

	fmt.Println(threshold, randNum)

	if randNum > threshold {
		cfg.totalExperience += 1
		fmt.Printf(" - You gained 1xp, total: %v\n", cfg.totalExperience)
		return fmt.Errorf("failed to catch %s", pokemonName)
	}

	cfg.caughtPokemon[pokemonName] = pokemon

	fmt.Printf("%s was caught!\n", pokemonName)
	cfg.totalExperience += 8
	fmt.Printf(" - You gained 3xp, total: %v\n", cfg.totalExperience)
	return nil
}
