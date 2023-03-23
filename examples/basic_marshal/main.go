package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func marshal(in any) []byte {
	out, err := json.Marshal(in)
	if err != nil {
		log.Fatalf("Unable to marshal due to %s\n", err)
	}

	return out
}

func main() {
	first := marshal(14)
	second := marshal("Hello world")
	third := marshal([]float32{1.66, 6.86, 10.1})
	fourth := marshal(map[string]int{"num": 15, "other": 17})
	fmt.Printf(
		"first: %s\nsecond: %s\nthird: %s\nfourth: %s\n",
		first,
		second,
		third,
		fourth,
	)
}
