package main

import "fmt"

func main() {

	ch := make(chan int)
	// cs := make(chan<- int)
	// defer close(cs)
	// cr := make(<-chan int)

	go func(cs chan<- int) {
		defer close(ch)
		for i := range 10 {
			cs <- i * 3
		}
	}(ch)

	fmt.Println("Recebendo valores do canal bidirecional...")
	for v := range ch {
		fmt.Println(v)
	}

}
