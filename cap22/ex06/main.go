package main

import "fmt"

func main() {

	ch := make(chan string)

	go func() {
		for range 100 {
			ch <- "http stream..."
		}
		close(ch)
	}()

	for s := range ch {
		fmt.Println("receiving", s)
	}

}
