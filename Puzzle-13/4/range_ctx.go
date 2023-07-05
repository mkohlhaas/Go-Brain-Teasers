package main

import (
	"context"
	"fmt"
)

// Dave Cheney: “Never start a goroutine without knowing how it will finish.”

func fibs(ctx context.Context, n int) chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		a, b := 1, 1
		for i := 0; i < n; i++ {
			select {
			case ch <- a:
				a, b = b, a+b
			case <-ctx.Done():
				fmt.Println("cancelled")
			}
		}
	}()

	return ch
}

// common practice is to pass a done channel or a context.Context

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	ch := fibs(ctx, 100)
	for i := 0; i < 10; i++ {
		val := <-ch
		fmt.Printf("%d ", val)
	}
	fmt.Println()
	cancel()
}
