package main

import (
	"fmt"
)

func callbackExperience(cfg *config, args ...string) error {
	fmt.Printf(" - Your total experience is: %d\n", cfg.totalExperience)

	return nil
}
