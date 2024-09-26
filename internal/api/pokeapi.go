package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Prodigy00/pokedexcli/internal/pokecache"
	"io"
	"net/http"
)

type Result struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
type PokeAPI struct {
	Count    int      `json:"count"`
	Next     *string  `json:"next"`
	Previous *string  `json:"previous"`
	Results  []Result `json:"results"`
	cache    pokecache.CacheAPI
}

func NewPokeAPI(cache pokecache.CacheAPI) *PokeAPI {
	return &PokeAPI{
		cache: cache,
	}
}

func (p *PokeAPI) GetLocationAreas(url *string) (PokeAPI, error) {
	if url == nil {
		return PokeAPI{}, errors.New("url is empty or invalid")
	}

	hit, exists := p.cache.Get(*url)
	if exists {
		var pke PokeAPI
		err := json.Unmarshal(hit, &pke)
		if err != nil {
			return PokeAPI{}, err
		}
	}

	res, err := http.Get(*url)
	if err != nil {
		return PokeAPI{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return PokeAPI{}, fmt.Errorf("failed to read response body: %w", err)
	}

	var result PokeAPI

	err = json.Unmarshal(body, &result)
	if err != nil {
		return PokeAPI{}, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	p.cache.Add(*url, body)

	return result, nil
}
