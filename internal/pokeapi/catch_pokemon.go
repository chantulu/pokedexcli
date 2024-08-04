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

func (c *Client) CatchPokemon(m string) (bool, error) {
	url := baseURL + "/pokemon/" + m
	if m == "" {
		return false, errors.New("no pokemon was provided")
	}

	if val, ok := c.cache.Get(url); ok {
		pokemonResp := RespPokemon{}
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return false, err
		}
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	pokemonResp := RespPokemon{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return false, err
	}

	attempt := attemptToCatch(pokemonResp.BaseExperience)
	
	c.cache.Add(url, dat)
	if attempt {
		return true,nil
	}
	return false, nil
}
