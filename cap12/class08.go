package cap12

import (
	"fmt"
)

func class08() {

	x := retornaFunc()
	y := x(3)

	fmt.Println(y)

	fmt.Println(x(3))

	fmt.Println(retornaFunc()(3))
}

func retornaFunc() func(int) int { // funcao que retorna funcao que retorna int
	return func(i int) int {
		return i * 10
	}
}
