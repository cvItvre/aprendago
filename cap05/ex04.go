package cap05

import (
	"fmt"
)

func ex04() {
	x := 10

	fmt.Printf("%d - %#b - %#x\n", x, x, x)

	y := x << 1

	fmt.Printf("%d - %#b - %#x", y, y, y)
}
