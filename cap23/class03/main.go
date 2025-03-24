package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	fmt.Println("------Exemplo 01------")
	_, err := os.Open("no-file.txt")
	if err != nil {
		// println works as usual, writing the error to stdout
		fmt.Println("err happened", err)
	}

	fmt.Println("------Exemplo 02------")
	if err != nil {
		// timestamp + error. outputs to std logger
		log.Println("err happened", err)
	}

	fmt.Println("------Exemplo 03------")
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	// setting the log output to a file
	log.SetOutput(f)

	_, err = os.Open("no-file.txt")
	if err != nil {
		log.Println("err happened", err)
	}

	fmt.Println("check the log.txt fule in the directory.")

	fmt.Println("------Exemplo 04------")
	defer foo()
	if err != nil {
		// println + panic -> panic stop the goroutine but allow defered functions to run
		// also, panic allow recover()
		log.Panicln(err)
	}

	fmt.Println("------Exemplo 05------")
	bar(err)
	recover() // see how to use it
	fmt.Println("the program continues after panic->recover...")

	fmt.Println("------Exemplo 06------")
	if err != nil {
		log.Fatalln(err)
	}
}

func foo() {
	fmt.Println("When os.Exit() is called, deffered functions don't run")
}

func bar(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
