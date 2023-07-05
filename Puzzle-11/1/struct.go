package main

import (
	"fmt"
	"io"
	"time"
)

// Log is a log message
type Log1 struct {
	Message   string
	time.Time // field with no name (a.k.a embedding) -> Log type has all the methods and fields that time.Time has including fmt.Stringer interface
}

type Log2 struct {
	Message string
	Time    time.Time
}

// You can also embed interfaces in Go.
// ReadWriter will have Read() and Write() functions.
type ReadWriter interface {
	io.Reader
	io.Writer
}

func main() {
	ts := time.Date(2009, 11, 10, 0, 0, 0, 0, time.UTC)

	log1 := Log1{"Hello", ts}
	fmt.Printf("%v\n", log1)
	fmt.Printf("%+v\n", log1)
	fmt.Printf("%#v\n", log1)

	fmt.Println()

	log2 := Log2{"Hello", ts}
	fmt.Printf("%v\n", log2)
	fmt.Printf("%+v\n", log2)
	fmt.Printf("%#v\n", log2)
}
