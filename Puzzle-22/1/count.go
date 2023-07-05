package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	var wg sync.WaitGroup

	for i := 0; i < 1_000_000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
      count++ // not thread-safe: race condition (and memory barrier)
		}()

	}
	wg.Wait()
	fmt.Println(count)
}
