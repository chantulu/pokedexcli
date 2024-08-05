package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"math/rand/v2"
	"net/http"
)

func attemptToCatch(expYield int) bool{
	rolled := rand.IntN(500)
	return rolled > expYield
}

func GetPokemon(c* Client, m string) (RespPokemon, error){
	url := baseURL + "/pokemon/" + m
	if m == "" {
		return RespPokemon{}, errors.New("no pokemon was provided")
	}

	if val, ok := c.cache.Get(url); ok {
		pokemonResp := RespPokemon{}
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return RespPokemon{}, err
		}
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemon{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespPokemon{}, err
	}

	pokemonResp := RespPokemon{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return RespPokemon{}, err
	}
	c.cache.Add(url, dat)
	return pokemonResp, nil
}

func (c *Client) CatchPokemon(m string) (bool, error) {
	pokemonResp, err := GetPokemon(c,m)
	if err != nil {
		return false, err
	}
	attempt := attemptToCatch(pokemonResp.BaseExperience)
	if attempt {
		return true,nil
	}
	return false, nil
}
