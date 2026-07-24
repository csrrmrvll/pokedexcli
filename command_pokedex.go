package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.caughtPokemon) == 0 {
		fmt.Println("You haven't caught any Pokemon yet.")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Println(" - ", pokemon.Name)
	}
	return nil
}
