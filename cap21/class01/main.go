package main

import "fmt"

func main() {

	// channel type int
	ch := make(chan int)

	go func() {
		ch <- 42
	}()

	fmt.Println(<-ch)

	// channel with a buffer
	ch2 := make(chan int, 1)
	ch2 <- 12
	fmt.Println(<-ch2)

}
