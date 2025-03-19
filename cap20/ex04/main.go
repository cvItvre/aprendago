package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	contador := 0
	numGoRoutines := 1000

	var m sync.Mutex

	wg.Add(numGoRoutines)

	for range numGoRoutines {
		go func() {
			m.Lock()
			aux := contador
			runtime.Gosched()
			aux++
			contador = aux
			fmt.Println("Contador:", contador)
			wg.Done()
			m.Unlock()
		}()
	}

	wg.Wait()

	fmt.Println("Valor final contador:", contador)

}
