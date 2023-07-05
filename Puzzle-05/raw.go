package main

import (
	"fmt"
)

func main() {
	s := `a\tb` // raw string
	// s := 'a\tb'
	// s := "a\tb"
	fmt.Println(s)
	fmt.Println("\u2122") // using Unicode code points
	// multiline string in an HTTP request
	request := `GET / HTTP/1.1
Host: www.353solutions.com
Connection: Close
`
	fmt.Println(request)
}
