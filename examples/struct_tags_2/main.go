package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	Username string   `json:"username"`
	Password string   `json:"-"`
	Email    string   `json:"email"`
	Hobbies  []string `json:"hobbies,omitempty"`
}

func main() {
	user := User{
		Username: "admin",
		Password: "root",
		Email:    "admin@email.com",
	}

	b, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		log.Fatalf("Unable to marshal due to %s\n", err)
	}

	fmt.Println(string(b))
}
