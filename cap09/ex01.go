package cap09

import (
	"fmt"
)

func ex01() {
	array := [5]int{}
	array[0], array[1], array[2], array[3], array[4] = 1, 2, 3, 4, 5

	for _, v := range array {
		fmt.Println(v)
	}

	fmt.Printf("%T", array)
}
