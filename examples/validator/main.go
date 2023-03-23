package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
)

type User struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email"    validate:"required,email"`
	Age      int    `json:"age"      validate:"required,min=18,max=99"`
}

func main() {
	input := `{
    "username": "johndoe",
    "email": "johndoe@emai",
    "age": -14
  }`

	var user User

	err := json.Unmarshal([]byte(input), &user)
	if err != nil {
		log.Fatalf("Unable to marshal JSON due to %s", err)
	}

	fmt.Printf("User before validation: %v\n", user)

	err = validator.New().Struct(user)
	if err != nil {
		log.Fatalf("Validation failed due to %v\n", err)
	}
}
