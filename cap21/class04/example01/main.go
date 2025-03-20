package main

import "fmt"

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	val := 10

	go func(cs chan<- int) {
		// defer close(cs)
		for i := range val {
			if i%2 == 0 {
				cs <- i
			}
		}
	}(ch1)

	go func(cs chan<- int) {
		// defer close(cs)
		for i := range val {
			if i%2 != 0 {
				cs <- i
			}
		}
	}(ch2)

	for range val {
		select {
		case num := <-ch1:
			fmt.Printf("Recebi um valor par do ch1: %v\n", num)
		case num := <-ch2:
			fmt.Printf("Recebi um valor impar do ch2: %v\n", num)
		}
	}

}
