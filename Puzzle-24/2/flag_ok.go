package main

import (
	"fmt"
)

type Flag int

func main() {
	var i any = 3
	f, ok := i.(Flag)
	if !ok {
		fmt.Println("not a Flag")
		return
	}
	fmt.Println(f)
}
