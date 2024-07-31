package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

var endpoints = map[string]string{
	"map": "https://pokeapi.co/api/v2/location/",
}

type MapResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func mapF(cfg *config) error {
	var full_url string
	if *cfg.nextLocationsURL == "" {
		full_url = endpoints["map"]
	} else {
		full_url = *cfg.nextLocationsURL
	}
	res, err := http.Get(full_url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	mapResponse := MapResponse{}
	err = json.Unmarshal(body, &mapResponse)
	if err != nil {
		log.Fatal(err)
	}
	*cfg.nextLocationsURL = mapResponse.Next
	*cfg.prevLocationsURL = fmt.Sprintf("%v", mapResponse.Previous)
	for _, city := range mapResponse.Results {
		fmt.Println(city.Name)
	}
	return nil
}
func mapB(cfg *config) error {
	if *cfg.prevLocationsURL == "" {
		return errors.New("no previous page available")
	}
	full_url := *cfg.prevLocationsURL
	res, err := http.Get(full_url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	mapResponse := MapResponse{}
	err = json.Unmarshal(body, &mapResponse)
	if err != nil {
		log.Fatal(err)
	}
	*cfg.nextLocationsURL = mapResponse.Next
	*cfg.prevLocationsURL = fmt.Sprintf("%v", mapResponse.Previous)
	for _, city := range mapResponse.Results {
		fmt.Println(city.Name)
	}
	return nil
}