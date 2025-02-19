package cap08

import (
	"fmt"
)

func class08() {

	slice1 := make([]int, 5, 10)

	for k := range slice1 {
		slice1[k] = k + 1
	}

	fmt.Println(slice1)

	slice2 := append(slice1[0:2], slice1[4:]...)
	fmt.Println(slice2)

	fmt.Println(slice1)

	// [1 2 3 4 5]
	// [1 2 5]
	// [1 2 5 4 5]

	// evitar usar essa abordagem;
	// usar o mesmo slice ou for loop para copiar os dados;

}
