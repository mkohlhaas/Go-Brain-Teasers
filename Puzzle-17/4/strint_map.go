package main

import (
	"fmt"
)

// https://dave.cheney.net/high-performance-go-workshop/gophercon-2019.html#using_byte_as_a_map_key

func main() {
	m := map[string]int{
		"hello": 3,
	}
	key := []byte{'h', 'e', 'l', 'l', 'o'} // you have a []byte but not a string
	val := m[string(key)]                  // specific compiler optimisation: no memory allocation
	fmt.Println(val)                       // 3
}
