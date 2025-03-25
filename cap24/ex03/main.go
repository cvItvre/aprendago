package main

import (
	"fmt"
	"log"
)

type erroEspecial struct {
}

func (e erroEspecial) Error() string {
	return fmt.Sprintln("um erro especial aconteceu...")
}

func main() {

	foo(erroEspecial{})

}

func foo(err error) {
	log.Print(err)
}
