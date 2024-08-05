package main

import (
	"errors"
	"fmt"

	"github.com/chantulu/pokedexcli/internal/pokeapi"
)

func inspectPokedex(cfg *config, pokemonName string) (pokeapi.Pokemon, error){
	pokemon, ok := cfg.pokeDex[pokemonName]
	if ok {
		return pokemon, nil
	}
	return pokeapi.Pokemon{}, errors.New("you have not caught that pokemon")
}

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("usage: catch <pokemon>")
	}
	pokemonName := args[0]
	caughtPokemon, err := inspectPokedex(cfg, pokemonName)
	if err != nil {
		return err
	}
	fmt.Println("Name: " + caughtPokemon.Name)
	fmt.Println("Height: " + fmt.Sprintf("%v", caughtPokemon.Height))
	fmt.Println("Weight: " + fmt.Sprintf("%v", caughtPokemon.Weight))
	fmt.Println("Stats: ")
	for _,stat := range caughtPokemon.Stats{
		fmt.Printf("-%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	return nil
}
