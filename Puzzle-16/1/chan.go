package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 2)
  fmt.Println(cap(ch))
	ch <- 1
	ch <- 2
	<-ch
	close(ch)
	a := <-ch
	b := <-ch
	fmt.Println(a, b)
}
