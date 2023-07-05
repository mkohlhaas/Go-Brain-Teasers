package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3}
	a = Append(a, 4)
	fmt.Println(a) // [1 2 3 4]
	b := a[:3]
	fmt.Println(b) // [1 2 3]
	b = append(b, 5)
	fmt.Println(a) // [1 2 3 5]
	fmt.Println(b) // [1 2 3 5]
}

// possible implementation of append
func Append(items []int, i int) []int {
	if len(items) == cap(items) { // No more space in underlying array
		// Go has a better growth heuristic than adding 1 every append
		newItems := make([]int, len(items)+1)
		copy(newItems, items)
		items = newItems
	} else {
		items = items[:len(items)+1]
	}

	items[len(items)-1] = i
	return items
}
