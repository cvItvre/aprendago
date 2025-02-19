package cap13

import (
	"fmt"
)

func ex10() {

	closure := criaSlice()

	closure()
	closure()
	closure()
	closure()
	closure()
	closure()

}

func criaSlice() func() {

	si := []int{0, 1, 2, 3, 4, 5}

	return func() {
		fmt.Println(si)
		si = append(si, len(si))
	}

}
