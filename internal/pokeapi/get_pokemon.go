package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokeName string) (t_Pokemon, error) {
	endpoint := fmt.Sprintf("/pokemon/%s", pokeName)
	fullURL := baseURL + endpoint

	if data, ok := c.cache.Get(fullURL); ok {
		out := t_Pokemon{}
		json.Unmarshal(data, &out)
		return out, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return t_Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return t_Pokemon{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return t_Pokemon{}, fmt.Errorf("Bad status code: %v", resp.StatusCode)
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return t_Pokemon{}, err
	}
	c.cache.Add(fullURL, dat)

	out := t_Pokemon{}
	json.Unmarshal(dat, &out)
	return out, nil
}
