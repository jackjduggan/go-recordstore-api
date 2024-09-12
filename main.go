package main

import (
	"encoding/json"
	"os"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums []album

func loadAlbumsFromJSON(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &albums)
	if err != nil {
		return err
	}
	return nil
}

func saveAlbumsToFile(filename string) error {
	file, err := json.MarshalIndent(albums, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, file, 0644)
	if err != nil {
		return err
	}
	return nil
}
