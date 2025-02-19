package cap12

import (
	"fmt"
)

func class06() {

	num := 10

	func(x int) { // funçoes anonimas
		fmt.Println(x, "* 10 = ", x*10)
	}(num)

}
