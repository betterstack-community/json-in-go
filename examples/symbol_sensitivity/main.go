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
      "Name": "Paws",
      "favorite_treat": "Kibble"
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
