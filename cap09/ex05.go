package cap09

import (
	"fmt"
)

func ex05() {

	x := []int{42, 43, 44, 45, 46,
		47, 48, 49, 50, 51}

	y := append(x[:3], x[len(x)-4:]...)

	fmt.Println(x, y)

}
