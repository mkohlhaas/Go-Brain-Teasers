package main

import (
	"fmt"
)

type Flag int

func main() {
	var i any = 3
	f := i.(Flag)
	fmt.Println(f)
}
