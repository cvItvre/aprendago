package cap18

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func class06() {

	// race condition

	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("GoRoutines:", runtime.NumGoroutine())

	var contador int64
	contador = 0
	totalGoRoutines := 1000

	var wg sync.WaitGroup
	wg.Add(totalGoRoutines)

	for i := 0; i < totalGoRoutines; i++ {
		go func() {
			atomic.AddInt64(&contador, 1)
			runtime.Gosched() // yield
			fmt.Println("Contador:\t", atomic.LoadInt64(&contador))

			wg.Done()
		}()
		fmt.Println("GoRoutines:", runtime.NumGoroutine())
	}

	wg.Wait()

	fmt.Println("GoRoutines:", runtime.NumGoroutine())
	fmt.Println("Valor Finall", contador)

}
