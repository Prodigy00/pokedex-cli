package api

import (
	"encoding/json"
	"errors"
	"fmt"
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
}

func GetLocationAreas(url *string) (PokeAPI, error) {
	validUrl := *url
	if validUrl == "" {
		return PokeAPI{}, errors.New("url is empty or invalid")
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

	return result, nil
}
