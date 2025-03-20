package main

import "fmt"

func main() {

	ch := make(chan int)

	go func() {
		ch <- 0
		close(ch)
	}()

	num, ok := <-ch

	fmt.Println(num, ok)

	num, ok = <-ch

	fmt.Println(num, ok)

}
