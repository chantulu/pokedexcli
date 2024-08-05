package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	
	fmt.Println("Your Pokedex:")
	if len(cfg.pokeDex) == 0 {
		fmt.Println("You currently have no pokemon")
		return nil
	}

	for _,pokemon := range cfg.pokeDex{
		fmt.Println(" - " + pokemon.Name)
	}
	
	return nil
}
