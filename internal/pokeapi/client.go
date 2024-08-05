package pokeapi

import (
	"net/http"
	"time"

	"github.com/chantulu/pokedexcli/internal/pokecache"
)

// Client -
type Client struct {
	cache pokecache.Cache
	httpClient http.Client
}

type Pokemon struct{
	Name string
	Height int `json:"height"`
	Weight int `json:"weight"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
}

// NewClient -
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: *pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}