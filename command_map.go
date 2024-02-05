package main

import (
	"fmt"

	"github.com/lucas-giraldelli/pokedexcli/internal/pokeapi"
)

func callbackMap() error {
  pokeApiClient := pokeapi.NewClient();

  resp, err := pokeApiClient.ListLocationAreas()

  if err != nil {
    return err
  }

  fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %v\n", area.Name)
	}

	return nil
}