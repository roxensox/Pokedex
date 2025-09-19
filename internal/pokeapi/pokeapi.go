package pokeapi

import (
	"github.com/roxensox/pokedexcli/internal/pokecache"
	"net/http"
	"time"
)

const baseURL = "https://pokeapi.co/api/v2/"

type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

func NewClient() Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
		cache: pokecache.NewCache(time.Minute * 5),
	}
}
