package cap06

import (
	"fmt"
)

func clas11() {

	x := 6

	if x == 2 || x == 3 {
		fmt.Println("x é dois ou tres")
	} else if x%2 == 0 && x%3 == 0 {
		fmt.Println("é múltiplo de 2 e 3")
	}

	// true
	// false
	// true
	// true
	// false

}
