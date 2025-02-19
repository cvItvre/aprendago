package cap12

import (
	"fmt"
)

func class07() {

	num := 2

	y := func(x int) int { // funçao como expressoes
		return x * 10
	}

	fmt.Println(y(num))

}
