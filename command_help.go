package main

import "fmt"

func callbackHelp() error {
  fmt.Println("Welcome to Pokedex help menu!")
  fmt.Println("Here are your available commands:")

  availableCommands := getCommands()

  for _, cmd := range availableCommands {
    fmt.Printf("- %v: %s\n",
      cmd.name,
      cmd.description,
    );
  }

  fmt.Println("")

  return nil
}

