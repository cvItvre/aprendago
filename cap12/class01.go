package cap12

import (
	"fmt"
)

func class01() {

	basica()
	argumento("pedro")
	fmt.Println(somaRetorno(5, 7))
	fmt.Println(variosRetornos(1, 5, 3, 6, 2, 7, 6))

}

func basica() {
	fmt.Println("oi, bom dia!")
}

func argumento(nome string) {
	fmt.Println("Oi,", nome, ", bom dia!")
}

func somaRetorno(x, y int) int { //x int, y, int
	return x + y
}

func variosRetornos(x ...int) (int, int) {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma, len(x)
}
