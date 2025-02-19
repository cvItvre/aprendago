package cap12

import (
	"fmt"
)

func class03() {

	defer fmt.Println("com defer") // o que vem primeiro, sai por ultimo (pilha)
	fmt.Println("sem defer")

	defer fmt.Println("1") // defer executa depois do escopo acabar
	fmt.Println("2")
	fmt.Println("3")
	fmt.Println("4")

}

// abrearquivo('filename')
// defer close() -> go design pattern
// usar os dados do arquivo
