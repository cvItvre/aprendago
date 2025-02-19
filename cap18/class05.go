package cap18

import (
	"fmt"
	"runtime"
	"sync"
)

func class05() {

	// race condition

	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("GoRoutines:", runtime.NumGoroutine())

	contador := 0
	totalGoRoutines := 1000

	var wg sync.WaitGroup
	wg.Add(totalGoRoutines)

	var mu sync.Mutex // tipo mutex

	for i := 0; i < totalGoRoutines; i++ {
		go func() {
			mu.Lock()
			aux := contador
			runtime.Gosched() // yield
			aux++
			contador = aux
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("GoRoutines:", runtime.NumGoroutine())
	}

	wg.Wait()

	fmt.Println("GoRoutines:", runtime.NumGoroutine())
	fmt.Println("Valor Finall", contador)

}
