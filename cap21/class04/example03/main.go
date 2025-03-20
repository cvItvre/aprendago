package main

import "fmt"

func main() {

	even := make(chan int)
	// defer close(even)
	odd := make(chan int)
	// defer close(odd)
	quit := make(chan int)
	// defer close(quit)

	go send(even, odd, quit)
	receive(even, odd, quit)

}

func send(even, odd, quit chan<- int) {
	for num := range 15 {
		if num%2 == 0 {
			even <- num
		} else {
			odd <- num
		}
	}
	// i can close the chan here after sending everthing
	close(even)
	close(odd)
	quit <- -1
}

func receive(even, odd, quit <-chan int) {
	for {
		select {
		case num, ok := <-even:
			if ok {
				fmt.Println("Receiving from chann even:", num)
			}
		case num, ok := <-odd:
			if ok {
				fmt.Println("Receiving from chann odd:", num)
			}
		case <-quit:
			return
		}
	}
}
