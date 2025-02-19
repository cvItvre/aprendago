package cap03

import (
	"fmt"
)

var a int = 42
var b string = "James Bond"
var c bool = true

func ex03() {

	s := fmt.Sprintf("x = %v \ny = %v \nz = %v", a, b, c)

	fmt.Println(s)

}
