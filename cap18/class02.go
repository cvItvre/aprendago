package cap18

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func class02() {

	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())

	// sync.WaitGroup{} -> func Add(): quantas go outines irei ter

	wg.Add(2)

	go func1()
	go func2()

	fmt.Println(runtime.NumGoroutine())

	wg.Wait()

	// sync.WaitGroup{} -> func Wait(): espera todo mundo terminar

}

func func1() {
	for i := 0; i < 1000; i++ {
		fmt.Println("func1:", i)
		time.Sleep(20) // sleep 20ms
	}
	// sync.WaitGroup{} -> func Done(): utilizada dentro de go routines
	// "Deu! terminei de executar" -> sinal para o waitgroup da funcao main
	wg.Done()
}

func func2() {
	for i := 0; i < 1000; i++ {
		fmt.Println("func2:", i)
		time.Sleep(20) // nao necessario
	}
	wg.Done()
}
