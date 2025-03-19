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

	wg.Add(numGoRoutines)

	for i := 0; i < numGoRoutines; i++ {
		go func() {
			aux := contador
			runtime.Gosched()
			aux++
			contador = aux
			wg.Done()
		}()
		fmt.Println("Contador:", contador)
	}

	wg.Wait()

	fmt.Println("Valor final contador:", contador)

}
