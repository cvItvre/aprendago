package cap09

import (
	"fmt"
)

func ex03() {
	slice := make([]int, 10, 20)

	for k, _ := range slice {
		slice[k] = (k + 1) * 10
	}

	fmt.Println(slice)
	fmt.Println(slice[:3])
	fmt.Println(slice[4:])
	fmt.Println(slice[1:7])
	fmt.Println(slice[2 : len(slice)-1])
}
