package main

import (
	"errors"
	"fmt"

	"github.com/chantulu/pokedexcli/internal/pokeapi"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("usage: catch <pokemon>")
	}
	pokemonName := args[0]
	catchResp, err := cfg.pokeapiClient.CatchPokemon(pokemonName)
	if err != nil {
		return err
	}
	fmt.Println("Throwing a Pokeball at " + pokemonName)
	if catchResp {
		fmt.Println(pokemonName + " was caught!")
		pokemon, err2 := pokeapi.GetPokemon(&cfg.pokeapiClient, pokemonName)
		if err2 != nil {
			return err2
		}
		cfg.pokeDex[pokemonName] = pokeapi.Pokemon{Name: pokemon.Name, Height: pokemon.Height, Weight: pokemon.Weight, Stats: pokemon.Stats}
		return nil
	}
	fmt.Println(pokemonName + " escaped!")
	return nil
}
