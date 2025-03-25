package escrevendo

import "fmt"

// x é um numero constante que eu nao uso.
const x = 100

// Do nothing...win
func Doc() {
	fmt.Println("do nothing...win")
}

// Foo func.
func foo() {
	fmt.Println("do nothing...win")
}

// Func bar é exportada.
func Bar() {
	fmt.Println("do nothing...win")
}
