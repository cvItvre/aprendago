package cap13

import (
	"fmt"
)

func ex07() {

	num := 25

	dobroNum := func(x int) int {
		return x * 2
	}(num)

	dobroNum2 := func(x int) int {
		return x * 2
	}

	fmt.Println("Dobro de", num, "é =", dobroNum)
	fmt.Println("Dobro de", num, "é =", dobroNum2(num))

}
