package cap07

import (
	"fmt"
)

func ex05() {
	for num := 10; num <= 100; num++ {
		fmt.Printf("%v %% 4 = %v\n", num, num%4)
	}
}
