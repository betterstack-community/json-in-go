package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Dog struct {
	Breed         string `json:"breed"`
	Name          string `json:"name"`
	FavoriteTreat string `json:"favorite_treat"`
	Age           int    `json:"age"`
}

func main() {
	newDog := Dog{
		Breed:         "Poodle",
		Age:           15,
		Name:          "Chloe",
		FavoriteTreat: "Dog Sticks",
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		encoder := json.NewEncoder(w)

		err := encoder.Encode(newDog)
		if err != nil {
			log.Fatalf("Unable to encode due to %s\n", err)
		}
	})

	log.Fatal(http.ListenAndServe(":3000", mux))
}
