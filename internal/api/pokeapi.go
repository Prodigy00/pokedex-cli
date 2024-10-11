package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Prodigy00/pokedexcli/internal/pokecache"
	"io"
	"net/http"
)

var (
	ErrNotFound = errors.New("not found")
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

func (p *PokeAPI) GetLocationArea(name string) ([]string, error) {
	if name == "" {
		return []string{}, errors.New("location name is empty or invalid")
	}
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", name)

	hit, exists := p.cache.Get(url)
	if exists {
		var locationArea LocationArea
		err := json.Unmarshal(hit, &locationArea)
		if err != nil {
			return []string{}, fmt.Errorf("failed to unmarshal locationArea: %w", err)
		}

		var pokemons []string
		for _, k := range locationArea.PokemonEncounters {
			pokemons = append(pokemons, k.Pokemon.Name)
		}
		return pokemons, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return []string{}, err
	}
	if res.StatusCode == http.StatusNotFound {
		return []string{}, ErrNotFound
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return []string{}, fmt.Errorf("failed to read response body: %w", err)
	}

	var result LocationArea

	err = json.Unmarshal(body, &result)
	if err != nil {
		return []string{}, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	var pokemons []string
	for _, k := range result.PokemonEncounters {
		pokemons = append(pokemons, k.Pokemon.Name)
	}

	p.cache.Add(url, body)

	return pokemons, nil
}

func (p *PokeAPI) CatchPokemon(name string) (CatchPokemonResult, error) {
	if name == "" {
		return CatchPokemonResult{}, errors.New("pokemon name is empty or invalid")
	}
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", name)

	hit, exists := p.cache.Get(url)
	if exists {
		var caughtPokemon CatchPokemonResult
		err := json.Unmarshal(hit, &caughtPokemon)
		if err != nil {
			return CatchPokemonResult{}, fmt.Errorf("failed to unmarshal caught_pokemon: %w", err)
		}
		return caughtPokemon, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return CatchPokemonResult{}, fmt.Errorf("failed to fetch pokemon: %w", err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return CatchPokemonResult{}, fmt.Errorf("failed to read response body: %w", err)
	}
	var result CatchPokemonResult
	err = json.Unmarshal(body, &result)
	if err != nil {
		return CatchPokemonResult{}, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	p.cache.Add(url, body)

	return result, nil
}
