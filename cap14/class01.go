package cap14

import (
	"fmt"
)

func class01() {

	x := 10
	y := &x

	fmt.Println(x)  // 10
	fmt.Println(&x) // 0xc000010070

	fmt.Println(y)  // 0xc000104040
	fmt.Println(&y) // 0xc000106038

	fmt.Println(x)   // 10
	fmt.Println(*y)  // 10  de-reference
	fmt.Println(*&x) // 10

	fmt.Printf("%T - %T\n", x, y) // int - *int

	a := 0
	b := &a
	fmt.Println(a) // 0

	*b++
	fmt.Println(a) // 1

}
