package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {

	var contador int64 = 0
	numGoRoutines := 1000

	wg.Add(numGoRoutines)

	for i := 0; i < numGoRoutines; i++ {
		go func() {
			atomic.AddInt64(&contador, 1)
			runtime.Gosched()
			fmt.Println("Contador:", atomic.LoadInt64(&contador))
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Valor final contador:", contador)

}
