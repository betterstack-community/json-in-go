package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/sanity-io/litter"
)

type (
	FullPerson struct {
		Address Address
		Name    string
		Pets    []Pet
		Age     int
	}

	Pet struct {
		Name  string
		Kind  string
		Color string
		Age   int
	}

	Address struct {
		Line1  string
		Line2  string
		Postal string
	}
)

func main() {
	b, err := os.ReadFile("assets/complex.json")
	if err != nil {
		log.Fatalf("Unable to read file due to %s\n", err)
	}

	var person FullPerson

	err = json.Unmarshal(b, &person)
	if err != nil {
		log.Fatalf("Unable to marshal JSON due to %s", err)
	}

	litter.Dump(person)
}
