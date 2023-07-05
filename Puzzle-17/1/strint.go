package main

import (
	"fmt"
)

func main() {
	i := 169
	s := string(i)
	fmt.Println(s)
	fmt.Println(fmt.Sprint(i)) // converting i of type int to a string
}
