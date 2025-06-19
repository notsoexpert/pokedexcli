package pokeapi

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/notsoexpert/pokedexcli/internal/pokecache"
)

type Location struct {
	Count     int                    `json:"count"`
	Base      string                 `json:"-"`
	Current   string                 `json:"-"`
	Next      *string                `json:"next"`
	Previous  *string                `json:"previous"`
	EndPoints []LocationAreaEndpoint `json:"results"`
}

var gLocalCache *pokecache.Cache

func init() {
	gLocalCache = pokecache.NewCache(5 * time.Second)
}

func unmarshalLocations(jsonData []byte, location *Location) error {
	if err := json.Unmarshal(jsonData, location); err != nil {
		return fmt.Errorf("error unmarshalling JSON: %w", err)
	}
	return nil
}

func UpdateLocations(url string, location *Location) error {
	body, ok := gLocalCache.Get(url)
	if !ok {
		var err error
		body, err = retrieveAPIData(url)
		if err != nil {
			return err
		}
		gLocalCache.Add(url, body)
	}

	err := unmarshalLocations(body, location)
	if err != nil {
		return err
	}
	return nil
}

func PrintLocations(location *Location) {
	for _, result := range location.EndPoints {
		fmt.Printf("%s\n", result.Name)
	}
}
