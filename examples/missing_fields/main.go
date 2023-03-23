package main

import (
	"encoding/json"
	"fmt"
	"log"
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
		"Name": "Paws"
	}`

	var dog Dog

	err := json.Unmarshal([]byte(input), &dog)
	if err != nil {
		log.Fatalf("Unable to marshal JSON due to %s", err)
	}

	fmt.Printf("%s likes %s\n", dog.Name, dog.FavoriteTreat)
}
