package main

import (
	"errors"
	"fmt"
)

func callbackMap(cfg *config) error {
  resp, err := cfg.pokeApiClient.ListLocationAreas(cfg.nextLocationArea)

  if err != nil {
    return err
  }

  fmt.Println("Location areas:")
  for _, area := range resp.Results {
    fmt.Printf(" - %v\n", area.Name)
  }

  cfg.nextLocationArea = resp.Next
  cfg.previousLocationArea = resp.Previous

  return nil
}

func callbackMapb(cfg *config) error {
  if cfg.previousLocationArea == nil {
    return errors.New("you're in the first page")
  }

  resp, err := cfg.pokeApiClient.ListLocationAreas(cfg.previousLocationArea)

  if err != nil {
    return err
  }

  fmt.Println("Location areas:")
  for _, area := range resp.Results {
    fmt.Printf(" - %v\n", area.Name)
  }

  cfg.nextLocationArea = resp.Next
  cfg.previousLocationArea = resp.Previous

  return nil
}