package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3}
  fmt.Printf("len: %v, cap: %v\n", len(a), cap(a))
  fmt.Printf("a=%v\n", a[:1])
	b := append(a[:1], 10) // changes underlying array
  fmt.Printf("len: %v, cap: %v\n", len(b), cap(b))
	fmt.Printf("a=%v, b=%v\n", a, b)
}
