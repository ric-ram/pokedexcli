package main

import (
	"fmt"
)

func callbackPokedex(cfg *config, args ...string) error {

	if len(cfg.caughtPokemon) == 0 {
		fmt.Println("Your Pokedex is currently empty")
		return nil
	}

	fmt.Println("Your Pokedex")
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}
