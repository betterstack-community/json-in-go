package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name    string
	Email   string
	Phone   string
	Hobbies []string
	Age     int
}

func main() {
	p := Person{
		Name:  "John Jones",
		Age:   26,
		Email: "johnjones@email.com",
		Phone: "89910119",
		Hobbies: []string{
			"Swimming",
			"Badminton",
		},
	}

	b, err := json.Marshal(p)
	if err != nil {
		log.Fatalf("Unable to marshal due to %s\n", err)
	}

	fmt.Println(string(b))
}
