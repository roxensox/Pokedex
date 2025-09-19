package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListEncounters(location string) (LocationDetails, error) {
	endpoint := fmt.Sprintf("/location-area/%s", location)
	fullURL := baseURL + endpoint

	if data, ok := c.cache.Get(fullURL); ok {
		out := LocationDetails{}
		json.Unmarshal(data, &out)
		return out, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)

	if err != nil {
		return LocationDetails{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationDetails{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		return LocationDetails{}, fmt.Errorf("Bad status code: %v", resp.StatusCode)
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationDetails{}, err
	}

	c.cache.Add(fullURL, dat)

	out := LocationDetails{}
	json.Unmarshal(dat, &out)
	return out, nil
}
