package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	rand.Seed(time.Now().UnixNano())
	catchChance := 0.5 * float64(pokemon.BaseExperience) // chance to catch the Pokemon
	if rand.Float64() > catchChance {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}
	fmt.Printf("%s was Caught!\n", pokemon.Name)
	return nil
}
