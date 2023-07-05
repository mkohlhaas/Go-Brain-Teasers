package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func main() {
	t1 := time.Now()
	data, err := json.Marshal(t1)  // when Go serializes a time.Time to JSON, it doesn't include the monotonic clock in the output
	if err != nil {
		log.Fatal(err)
	}

	var t2 time.Time
	if err := json.Unmarshal(data, &t2); err != nil {
		log.Fatal(err)
	}

  fmt.Println(t1)
  fmt.Println(t2) // monotonic clock is lost

	fmt.Println(t1 == t2)
	fmt.Println(t1.Equal(t2)) // correct way to do it
}
