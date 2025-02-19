package cap06

import "fmt"

func class04() {

	x := 0

	for x < 10 {
		fmt.Println(x)
		x++
	}

	for {
		if x < 20 {
			x++
			fmt.Println("x é menor que 20")
		} else {
			break
		}
	}

}
