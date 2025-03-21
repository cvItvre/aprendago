package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	ch := converge(work("channel 1"), work("channel 2"))

	for range 50 {
		fmt.Printf("receiving from %s...\n", <-ch)
	}

}

func work(s string) chan string {

	ch := make(chan string)

	go func(c chan string, s string) {
		for {
			c <- s
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
		}
	}(ch, s)

	return ch

}

func converge(ch1, ch2 chan string) chan string {

	ch := make(chan string)

	go func(s, r chan string) {
		for {
			r <- <-s
		}
	}(ch1, ch)

	go func(s, r chan string) {
		for {
			r <- <-s
		}
	}(ch2, ch)

	return ch

}
