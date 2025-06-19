package pokeapi

import (
	"encoding/json"
	"fmt"
)

type LocationAreaEndpoint struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	GameIndex            int    `json:"game_index"`
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
			MaxChance        int `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel        int   `json:"min_level"`
				MaxLevel        int   `json:"max_level"`
				ConditionValues []any `json:"condition_values"`
				Chance          int   `json:"chance"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func unmarshalLocationAreaEndPoint(jsonData []byte, endPoint *LocationAreaEndpoint) error {
	if err := json.Unmarshal(jsonData, endPoint); err != nil {
		return fmt.Errorf("error unmarshalling JSON: %w", err)
	}
	return nil
}

func updateLocationAreaEndpoint(url string, endPoint *LocationAreaEndpoint) error {
	body, ok := gLocalCache.Get(url)
	if !ok {
		var err error
		body, err = retrieveAPIData(url)
		if err != nil {
			return err
		}
		gLocalCache.Add(url, body)
	}

	err := unmarshalLocationAreaEndPoint(body, endPoint)
	if err != nil {
		return err
	}
	return nil
}

func PrintExploration(area string, location *Location) {
	fmt.Printf("Exploring %s...\n", area)
	var endPoint LocationAreaEndpoint

	err := updateLocationAreaEndpoint(location.Base+area, &endPoint)
	if err != nil {
		fmt.Printf("Failed to retrieve data! Reason: %s\n", err.Error())
		return
	}
	for _, enc := range endPoint.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}
}
