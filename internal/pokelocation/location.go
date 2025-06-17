package pokelocation

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/notsoexpert/pokedexcli/internal/pokecache"
)

type Location struct {
	Count     int        `json:"count"`
	Current   string     `json:"-"`
	Next      *string    `json:"next"`
	Previous  *string    `json:"previous"`
	EndPoints []EndPoint `json:"results"`
}

type EndPoint struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

var gLocalCache *pokecache.Cache

func init() {
	gLocalCache = pokecache.NewCache(5 * time.Second)
}

func retrieveLocationData(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
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
		body, err = retrieveLocationData(url)
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
