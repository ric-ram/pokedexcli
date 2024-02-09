package main

import (
	"errors"
	"fmt"
)

func callbackExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no location area provided")
	}
	locationsAreaName := args[0]

	locationArea, err := cfg.pokeapiClient.GetLocationAreas(locationsAreaName)
	if err != nil {
		return err
	}
	fmt.Printf("Pokemon in %s:\n", locationsAreaName)
	for _, pokemon := range locationArea.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
