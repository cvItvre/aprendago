package cap06

import (
	"fmt"
)

func class08() {

	x := 50

	if x > 100 {
		fmt.Println("x é maior que 100")
	} else if x < 10 {
		fmt.Println("x é menor que 10")
	} else {
		fmt.Println("x está entre 10 e 100")
	}

}
