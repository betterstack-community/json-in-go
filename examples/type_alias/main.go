package main

import (
	"encoding/json"
	"log"

	"github.com/sanity-io/litter"
)

type (
	TypeAliasExample string

	TypeAliasStruct struct {
		Example TypeAliasExample
	}
)

func main() {
	input := `{"Example": "Hello world"}`

	var tas TypeAliasStruct

	err := json.Unmarshal([]byte(input), &tas)
	if err != nil {
		log.Fatalf("Unable to marshal JSON due to %s", err)
	}

	litter.Dump(tas)
}
