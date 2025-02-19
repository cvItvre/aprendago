package cap08

import (
	"fmt"
)

var x [5]int
var y [6]int

func class01() {
	x[0] = 1
	x[1] = 10
	fmt.Println(x[0], x[1], x)
	fmt.Printf("%T\n", y)
	fmt.Println(len(x))
}
