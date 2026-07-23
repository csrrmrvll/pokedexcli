package main

import (
	"fmt"

	"github.com/csrrmrvll/pokedexcli/internal"
)

func _getLocationAreas(start, end int) error {
	// fmt.Println("start:", start, "end:", end)
	for i := start; i < end; i++ {
		location, err := internal.GetLocationArea(i)
		if err != nil {
			return err
		}
		fmt.Printf("%s\n", location)
	}
	return nil
}

func commandMap(config *config) error {
	err := _getLocationAreas(config.next, config.next+20)
	if err != nil {
		return err
	}
	config.next += 20
	return nil
}

func commandMapB(config *config) error {
	err := _getLocationAreas(config.previous, config.previous+20)
	if err != nil {
		return err
	}
	config.previous -= 20
	if config.previous < 1 {
		config.previous = 1
	}
	config.next = config.previous + 20
	return nil
}
