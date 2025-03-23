package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go send(ch1, 1000000)
	go receiveAndCreate(ch1, ch2)

	for num := range ch2 {
		fmt.Println("receiving", num, "...")
	}

}

func send(ch chan int, num int) {
	for i := range num {
		ch <- i
	}
	close(ch)
}

func receiveAndCreate(cs, cr chan int) {
	var wg sync.WaitGroup
	for range cs {
		wg.Add(1)
		go func() {
			cr <- work()
			wg.Done()
		}()
	}
	wg.Wait()
	close(cr)
}

func work() int {
	val := rand.Intn(1e3)
	time.Sleep(time.Millisecond * time.Duration(val))
	return val
}
