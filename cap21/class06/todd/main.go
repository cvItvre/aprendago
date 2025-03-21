package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	even := make(chan int)
	odd := make(chan int)
	converge := make(chan int)

	go send(even, odd)
	go receive(even, odd, converge)

	for num := range converge {
		fmt.Println("receiving...", num)
	}

}

func send(even, odd chan<- int) {

	defer close(even)
	defer close(odd)

	for num := range 100 {
		if num%2 == 0 {
			even <- num
		} else {
			odd <- num
		}
	}

}

func receive(even, odd <-chan int, converge chan<- int) {

	wg.Add(2)

	go func() {
		for num := range even {
			converge <- num
		}
		wg.Done()
	}()

	go func() {
		for num := range odd {
			converge <- num
		}
		wg.Done()
	}()

	wg.Wait()

	close(converge)

}
