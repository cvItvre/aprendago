package cap12

import (
	"fmt"
)

func class11() {

	fmt.Printf("Fatorial de %v = %v\n", 5, fatorial(5))
	fmt.Printf("Fatorial de %v = %v", 5, fatorial2(5))

}

func fatorial(num int) int {
	if num == 1 {
		return 1
	} else {
		return num * fatorial(num-1)
	}
}

func fatorial2(num int) int {
	resultado := 1
	for num > 1 {
		resultado *= num
		num--
	}
	return resultado
}
