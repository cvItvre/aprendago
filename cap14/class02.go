package cap14

import (
	"fmt"
)

func class02() {

	x := 0

	recebeValor(x)
	fmt.Println(x) // 0

	recebePointer(&x)
	fmt.Println(x) // 1

}

func recebeValor(x int) {
	x++
	fmt.Println("Nao funcao:", x) // 1
}

func recebePointer(x *int) {
	*x++
	fmt.Println("Nao funcao:", *x) // 1
}
