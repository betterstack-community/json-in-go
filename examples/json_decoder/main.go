package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/sanity-io/litter"
)

type Dog struct {
	Breed         string `json:"breed"`
	Name          string `json:"name"`
	FavoriteTreat string `json:"favorite_treat"`
	Age           int    `json:"age"`
}

func main() {
	coffeeFile, err := os.Open("assets/coffee.json")
	if err != nil {
		log.Fatalf("Unable to read file due to %s\n", err)
	}

	var coffee Dog

	decoder := json.NewDecoder(coffeeFile)
	decoder.DisallowUnknownFields()

	err = decoder.Decode(&coffee)
	if err != nil {
		log.Fatalf("Unable to decode due to %s\n", err)
	}

	litter.Dump(coffee)
}
