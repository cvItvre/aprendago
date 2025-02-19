package cap08

import (
	"fmt"
)

func class05() {
	umaslice := []int{1, 2, 3, 4, 5}
	outraslice := []int{6, 7, 8, 9}

	umaslice = append(umaslice, outraslice...)
	fmt.Println(umaslice)

	umaslice = append(umaslice, 10, 11, 12, 13, 14, 15)
	fmt.Println(umaslice)
}
