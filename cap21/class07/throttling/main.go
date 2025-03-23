package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go send(ch1, 100)
	go receiveAndCreate(ch1, ch2, 5)

	for num := range ch2 {
		fmt.Println("receiving...", num)
	}

}

func send(ch chan int, num int) {
	for i := range num {
		ch <- i
	}
	close(ch)
}

func receiveAndCreate(cs, cr chan int, threads int) {
	var wg sync.WaitGroup
	wg.Add(threads)
	for range threads {
		go func() {
			for num := range cs {
				cr <- work(num)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(cr)
}

func work(x int) int {
	// val := rand.Intn(1e3)
	time.Sleep(time.Millisecond * 1000)
	return x
}
