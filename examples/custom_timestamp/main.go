package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/araddon/dateparse"
)

type (
	CustomTime struct {
		time.Time
	}

	Baby struct {
		BirthDate CustomTime `json:"birth_date"`
		Name      string     `json:"name"`
		Gender    string     `json:"gender"`
	}
)

func (ct *CustomTime) UnmarshalJSON(input []byte) error {
	value := strings.Trim(string(input), `"`)

	t, err := dateparse.ParseAny(value)
	if err != nil {
		return err
	}

	ct.Time = t

	return nil
}

func (ct CustomTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`%q`, ct.Time.Format("02-01-2006"))), nil
}

func main() {
	var baby Baby

	input := []byte(
		`{"name": "Mary", "gender": "F", "birth_date": "2022/12/31"}`,
	)
	if err := json.Unmarshal(input, &baby); err != nil {
		log.Fatalf("Unable to unmarshal due to %s\n", err)
	}

	b, err := json.Marshal(baby)
	if err != nil {
		log.Fatalf("Unable to marshal due to %s\n", err)
	}

	fmt.Println(string(b))
}
