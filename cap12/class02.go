package cap12

import (
	"fmt"
)

func class02() {

	numeros := []int{1, 5, 3, 6, 2, 7, 6}
	fmt.Println(some(numeros...)) // slice unfurl
	fmt.Println(some())           // no arguments

}

func some(numeros ...int) int {
	soma := 0
	for _, v := range numeros {
		soma += v
	}
	return soma
}
