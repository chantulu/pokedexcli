package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("usage: explore <location>")
	}
	location := args[0]
	exploreResp, err := cfg.pokeapiClient.ExploreMap(location)
	if err != nil {
		return err
	}
	for _, name := range exploreResp.Names {
		fmt.Println(name.Name)
	}
	return nil
}
