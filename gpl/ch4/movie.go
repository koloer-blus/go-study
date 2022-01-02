package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{
		Title:  "Casablanca",
		Year:   1942,
		Color:  false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"},
	},
	{
		Title:  "Cool Hand Luke",
		Year:   1967,
		Color:  true,
		Actors: []string{"Paul Newman"},
	},
	{
		Title:  "Bullitt",
		Year:   1968,
		Color:  true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
}

func main() {
	data1, err := json.Marshal(movies)
	data2, _ := json.MarshalIndent(movies, "", "    ")
	if err != nil {
		log.Fatalf("JSON failed: %s", err)
	}
	fmt.Printf("%s\n%s", data1, data2)
}
