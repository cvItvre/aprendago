package main

import (
	"fmt"
)

func main() {

	ch := make(chan int)
	quit := make(chan int)

	val := 1

	go func(cr <-chan int, t int, cq chan int) {
		for range t {
			fmt.Println("Recebendo valor do ch:", <-cr)
		}
		cq <- 666
	}(ch, 15, quit)

	func(cs chan<- int, num int) {
		for {
			select {
			case ch <- num:
				fmt.Printf("Enviando valor %v para ch...\n", num)
				num++
			case num := <-quit:
				fmt.Println("Recebi valor do quit:", num)
				return
			}
		}
	}(ch, val)

}
