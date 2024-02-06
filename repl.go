package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(config *config) {
  scanner := bufio.NewScanner(os.Stdin);

  for {
    fmt.Print(" > ")
    scanner.Scan();
    text := scanner.Text()

    cleaned := cleanInput(text)

    if len(cleaned) == 0 {
      continue
    }

    commandName := cleaned[0]

    availableCommands := getCommands()

    command, ok := availableCommands[commandName]

    if !ok {
      fmt.Println("invalid command")
      continue
    }

    err := command.callback(config)
    if err != nil {
      fmt.Println(err)
    }
  }
}

type cliCommand struct {
  name        string
  description string
  callback    func(*config) error
}

func getCommands() map[string]cliCommand {
  return map[string]cliCommand{
    "help": {
      name: "help",
      description: "Prints the help menu",
      callback: callbackHelp,
    },
    "map": {
      name: "map",
      description: "Show pokemon locations areas",
      callback: callbackMap,
    },
    "mapb": {
      name: "mapb",
      description: "Show previous pokemon locations areas",
      callback: callbackMapb,
    },
    "exit": {
      name: "exit",
      description: "Turns off the Pokedex",
      callback: callbackExit,
    },
  }
}

func cleanInput(str string) []string {
  lower := strings.ToLower(str)
  words := strings.Fields(lower)

  return words
}
