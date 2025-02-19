package cap12

import (
	"fmt"
)

func class09() {

	total := somentePares(soma, []int{12, 2, 3, 5, 6, 32, 6, 43, 15, 7, 6, 23, 7, 34}...)
	fmt.Println("A soma de todos os numeros pares é:", total)

}

func soma(numeros ...int) int {
	soma := 0
	for _, numero := range numeros {
		soma += numero
	}
	return soma
}

// callback = func coo argumento
func somentePares(f func(collection ...int) int, numeros ...int) int {
	pares := []int{}
	for _, numero := range numeros {
		if numero%2 == 0 {
			pares = append(pares, numero)
		} else {
			continue
		}
	}
	return f(pares...)
}
