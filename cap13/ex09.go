package cap13

import (
	"fmt"
)

func ex09() {

	numeros := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(numeros)

	numeros = applyToSlice(dobrar, numeros)
	fmt.Println(numeros)

}

func dobrar(num int) int {
	return num * 2
}

func applyToSlice(f func(x int) int, numeros []int) []int {
	newSlice := []int{}
	for _, num := range numeros {
		newSlice = append(newSlice, f(num))
	}
	return newSlice
}
