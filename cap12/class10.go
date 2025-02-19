package cap12

import (
	"fmt"
)

func class10() {

	a := i()

	fmt.Println(a()) // x == 1
	fmt.Println(a()) // x == 2
	fmt.Println(a()) // x == 3

	b := i()

	fmt.Println(b()) // x == 1
	fmt.Println(b()) // x == 2
	fmt.Println(b()) // x == 3

}

func i() func() int {

	x := 0
	return func() int { // closure - capturamos o escopo e usamos quando quiser.
		x++
		return x
	}

}
