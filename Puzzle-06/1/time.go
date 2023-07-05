package main

import (
	"fmt"
	"time"
)

// const timeout = 3

func main() {
	timeout := 3
	// var timeout time.Duration = 3
	fmt.Printf("before ")
	time.Sleep(timeout * time.Second)
	fmt.Println("after")
}
