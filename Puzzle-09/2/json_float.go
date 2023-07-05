package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	n1 := 1
	data, err := json.Marshal(n1)
	if err != nil {
		log.Fatal(err)
	}

	var n2 any
	if err := json.Unmarshal(data, &n2); err != nil {
		log.Fatal(err)
	}

	// JSON only has floating-point numbers
	fmt.Println(n1 == n2)
	fmt.Printf("n1 is %T, n2 is %T\n", n1, n2)
}
