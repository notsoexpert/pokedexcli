package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"io"
)

type Location struct {
	Count    int    		`json:"count"`
	Current  string			`json:"-"`
	Next     *string 		`json:"next"`
	Previous *string 		`json:"previous"`
	EndPoints []EndPoint	`json:"results"`
}	

type EndPoint struct {
	Name string 	`json:"name"`
	URL  string 	`json:"url"`
}

func retrieveLocationData(url string, location *Location) ([]byte, error) {
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
        return fmt.Errorf("Error unmarshalling JSON: %w", err)
    }
    return nil
}

func updateLocations(url string, location *Location) error {
	var body []byte
	body, err := retrieveLocationData(url, location)
	if err != nil {
		return err
	}

	err = unmarshalLocations(body, location)
	if err != nil {
		return err
	}
	return nil
}

func printLocations(location *Location) {
	for _, result := range location.EndPoints {
		fmt.Printf("%s\n", result.Name)
	}
}