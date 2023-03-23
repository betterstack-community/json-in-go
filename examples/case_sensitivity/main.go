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
		"BreED": "Golden Retriever", 
		"age": 8, 
		"NaME": "Paws", 
		"favoriTeTrEat": "Kibble"
	}`

	var dog Dog

	err := json.Unmarshal([]byte(input), &dog)
	if err != nil {
		log.Fatalf("Unable to marshal JSON due to %s", err)
	}

	fmt.Printf(
		"%s is a %d years old %s who likes %s\n",
		dog.Name,
		dog.Age,
		dog.Breed,
		dog.FavoriteTreat,
	)
}
