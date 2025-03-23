package main

import (
	"fmt"
	"sync"
)

func main() {

	ch := make(chan int)

	go releaseWorkers(10, ch)

	for data := range ch {
		fmt.Println("receiving...", data)
	}

}

func releaseWorkers(numWorkers int, ch chan int) {
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for worker := range numWorkers {
		go func() {
			profile := worker
			for range 10 {
				ch <- profile
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(ch)
}
