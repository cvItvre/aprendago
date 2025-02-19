package cap13

import (
	"fmt"
)

func ex02() {

	numeros := []int{2, 4, 6, 13, 27, 9, 14, 3, 8, 1, 21}

	fmt.Println("Usando a funcao somaNumeros:", somaNumeros(numeros...))
	fmt.Println("Usando a funcao somaSlice:", somaSlice(numeros))

}

func somaNumeros(numeros ...int) int {
	soma := 0
	for _, numero := range numeros {
		soma += numero
	}
	return soma
}

func somaSlice(numeros []int) int {
	soma := 0
	for _, numero := range numeros {
		soma += numero
	}
	return soma
}
