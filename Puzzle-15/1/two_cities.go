package main

import (
	"fmt"
)

func main() {
	// city1 has a 1-byte rune at position 4 (ó), while city2 has the o rune in position 4
	// and a control character saying “add an umlaut to the previous character.”
	// These ways of encoding are called NFC and NFD.
	city1, city2 := "Kraków", "Kraków"

  fmt.Printf("len citi1: %v\nlen citi2: %v\n", len(city1), len(city2))
	fmt.Println(city1 == city2)
}
