package main

import (
	"fmt"
)

func commandExplore(cfg *config) error {
	areaDetailResp, err := cfg.pokeapiClient.LocationDetail(cfg.areaName)
	if err != nil {
		return err
	}

	fmt.Println("Exploring " + cfg.areaName + "...")
	fmt.Println("Found Pokemon:")
	for _, encounter := range areaDetailResp.PokemonEncounters {
		fmt.Println(" - " + encounter.Pokemon.Name)
	}
	return nil
}
