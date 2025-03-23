package main

import (
	"fmt"
)

func main() {

	// with auto-exec anomymous function
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)

	// with buffer
	ch := make(chan int, 1)

	ch <- 42

	fmt.Println(<-ch)

}
