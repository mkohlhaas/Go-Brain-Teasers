package main

import (
	"fmt"
)

type OSError int

func (e *OSError) Error() string {
	return fmt.Sprintf("error #%d", *e)
}

func FileExists(path string) (bool, error) {
	var err *OSError // pointer is not an interface
	// var err error // correct way to do it
	return false, err // TODO
}

func main() {
	if _, err := FileExists("/no/such/file"); err != nil {
		fmt.Printf("my bad: %s\n", err)
	} else {
		fmt.Println("OK")
	}
}
