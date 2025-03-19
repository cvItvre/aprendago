package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var totalGoRoutines int = 100

func main() {

	wg.Add(totalGoRoutines)

	for i := 1; i <= totalGoRoutines; i++ {
		go imprimir(i)
	}

	wg.Wait()

}

func imprimir(goroutine int) {
	fmt.Println("Olá, sou a go routine", goroutine)
	wg.Done()
}
