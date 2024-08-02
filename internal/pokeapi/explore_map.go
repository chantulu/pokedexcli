package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) ExploreMap(m string) (RespExploreMap, error) {
	url := baseURL + "/location-area/" + m
	if m == "" {
		return RespExploreMap{}, errors.New("no location provided")
	}

	if val, ok := c.cache.Get(url); ok {
		exploreResp := RespExploreMap{}
		err := json.Unmarshal(val, &exploreResp)
		if err != nil {
			return RespExploreMap{}, err
		}
		return exploreResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespExploreMap{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespExploreMap{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespExploreMap{}, err
	}

	exploreResp := RespExploreMap{}
	err = json.Unmarshal(dat, &exploreResp)
	if err != nil {
		return RespExploreMap{}, err
	}

	c.cache.Add(url, dat)
	return exploreResp, nil
}
