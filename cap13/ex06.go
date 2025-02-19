package cap13

import (
	"fmt"
)

func ex06() {

	num := 5

	func(x int) {
		fmt.Println("Dobro de", x, "é =", x*2)
	}(num)

}
