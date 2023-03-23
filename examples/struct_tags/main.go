package main

import (
	"encoding/json"
	"log"

	"github.com/sanity-io/litter"
)

type Dog struct {
	Breed         string `json:"breed"`
	Name          string `json:"name"`
	FavoriteTreat string `json:"favorite_treat"`
	Age           int    `json:"age"`
}

func main() {
	input := `{
  "name": "Coffee",
  "breed": "Toy Poodle",
  "age": 5,
  "favorite_treat": "Kibble"
}`

	var coffee Dog

	err := json.Unmarshal([]byte(input), &coffee)
	if err != nil {
		log.Fatalf("Unable to marshal JSON due to %s", err)
	}

	litter.Dump(coffee)
}
