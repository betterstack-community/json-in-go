package main

import (
	"encoding/json"
	"log"

	"github.com/sanity-io/litter"
)

type Dog struct {
	Breed         string
	Name          string
	FavoriteTreat string
	Age           int
}

func main() {
	input := `{
		"Breed": "Golden Retriever", 
		"Age": 8, 
		"Name": "Paws", 
		"FavoriteTreat": "Kibble", 
		"Dislikes": "Cats"
	}`

	var dog Dog

	err := json.Unmarshal([]byte(input), &dog)
	if err != nil {
		log.Fatalf("Unable to marshal JSON due to %s", err)
	}

	litter.Dump(dog)
}
