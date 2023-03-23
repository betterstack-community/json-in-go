package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	good := `{"name": "John Doe"}`
	bad := `{name: "John Doe"}`

	fmt.Println(json.Valid([]byte(good)))
	fmt.Println(json.Valid([]byte(bad)))
}
